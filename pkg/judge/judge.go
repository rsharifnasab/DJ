package judge

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/rsharifnasab/DJ/pkg/run"
	"github.com/rsharifnasab/DJ/pkg/ts"
	"github.com/rsharifnasab/DJ/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (submission *Submission) checkReq() {
	if viper.GetBool("debug") {
		abs, err := filepath.Abs(submission.sandboxDir)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("sandbox dir: %v\n", abs)
		}
	}
	stdout, err := run.JustOut(submission.sandboxDir + "/req.sh")
	cobra.CheckErr(err)
	space := regexp.MustCompile(`\s+`)
	reqWords := space.Split(stdout, -1)
	run.CheckAndErrorRequirements(reqWords)
}

func runTestCase(submission *Submission, i int) (result *TestResult) {

	command := fmt.Sprintf("%s/run.sh test %s %d", submission.sandboxDir, submission.Language, i)
	stdout, stderr, err := run.DefaultRun(command)
	//stdout, stderr, err := run.Run(command, 10*1024*1024, 1024*1024*1024, 4*time.Second)

	// TODO: write stderr to file
	submission.logger.LogTo(submission.currentGroup, strconv.Itoa(i), stderr)

	result = &TestResult{
		Run: true,
	}
	if err != nil {
		switch err {
		case run.ErrMemoryLimit:
			result.Killed = true
			return
		case run.ErrOutputLimit:
			result.Killed = true
			return
		case run.ErrMalformedCommand:
			panic(err)
		case run.ErrTimedOut:
			result.TimedOut = true
			return
		case run.ErrNoOutput:
			result.NoResult = true
			return
		case run.ErrNotValidExecutable:
			panic("cannot run test " + strconv.Itoa(i))
		case run.ErrNonZeroExit:
			result.NonZero = true
			return
		}
	} else {
		if viper.GetBool("debug") {
			fmt.Println("no error")
		}
	}
	var n int
	var resultStr string
	_, err = fmt.Sscanf(stdout, "test[%d] - %s\n", &n, &resultStr)
	if err != nil {
		result.MalformedOutput = true
		return
	}

	switch resultStr {
	case "pass":
		result.Pass = true
		return
	case "fail":
		result.Wrong = true
		return
	default:
		result.MalformedOutput = true
		return
	}

}

func (submission *Submission) currentGroupTestCount() int {
	command := fmt.Sprintf("%s/run.sh count %s",
		submission.sandboxDir, submission.Language)
	//println(command)
	stdout, err := run.JustOut(command)
	cobra.CheckErr(err)
	n, err := strconv.Atoi(strings.TrimSpace(stdout))
	cobra.CheckErr(err)
	return n
}

func (submission *Submission) compile() {
	command := fmt.Sprintf("%s/run.sh compile %s",
		submission.sandboxDir, submission.Language)

	stdout, stderr, err := run.DefaultRun(command)
	submission.logger.LogTo("", "compile", stdout)
	submission.logger.LogTo("", "compile", stderr)
	if err != nil && err != run.ErrNoOutput {
		cobra.CheckErr(fmt.Errorf("compilation failed:\nerr: %w", err))
	} else {
		fmt.Println("Compilation successful")
	}
}

func (submission *Submission) initSandboxWithoutTest() {
	util.CopyDir(submission.Judge, submission.sandboxDir)
	util.CopyDir(submission.solution, submission.sandboxDir+"/src/")
}

func (submission *Submission) backupCompiled() {
	util.CopyDir(submission.sandboxDir, submission.CompiledState)
}

func (submission *Submission) restoreCompiled() {
	err := os.RemoveAll(submission.sandboxDir)
	cobra.CheckErr(err)
	util.CopyDir(submission.CompiledState, submission.sandboxDir)
}

func (submission *Submission) exploreTestGroups() []*TestGroupResult {
	res := make([]*TestGroupResult, 0, 10)

	files, err := os.ReadDir(submission.Question)
	cobra.CheckErr(err)
	for _, f := range files {
		if f.IsDir() {
			testGroup := &TestGroupResult{
				Name:        f.Name(), // just name
				TestResults: make([]*TestResult, 0, 10),
			}
			res = append(res, testGroup)
		}
	}
	return res
}

func (submission *Submission) prepareForTestGroup(group *TestGroupResult) {
	src := fmt.Sprintf("%s/%s", submission.Question, group.Name)
	dest := fmt.Sprintf("%s/testgroup", submission.sandboxDir)
	util.CopyDir(src, dest)
	submission.currentGroup = group.Name // for logging
}

func (submission *Submission) createZipResult() {
	util.CopyDir(submission.solution, submission.ResultDir+"/solution")
	submission.Result.DumpTo(submission.ResultDir + "/result.txt")

	util.ZipDir(
		submission.ResultDir+"/submission.zip",
		submission.ResultDir,
	)
}

func (submission *Submission) PrintInitialInfo() {
	fmt.Printf("result dir: %v\n", submission.ResultDir)
	if viper.GetBool("debug") {
		fmt.Printf("Sandbox  dir: %s\n", submission.sandboxDir)
		fmt.Printf("compiled dir: %s\n", submission.CompiledState)
	} else {
		defer os.RemoveAll(submission.sandboxDir)
		defer os.RemoveAll(submission.CompiledState)
	}
}

func (submission *Submission) RunGroup(groupResult *TestGroupResult) {
	submission.restoreCompiled()
	submission.prepareForTestGroup(groupResult)
	println("-> " + groupResult.Name)

	groupResult.TestCount = submission.currentGroupTestCount()
	for i := 1; i <= groupResult.TestCount; i++ {
		singleTestResult := runTestCase(submission, i)
		groupResult.TestResults = append(groupResult.TestResults, singleTestResult)
	}

	if viper.GetBool("debug") {
		fmt.Println(groupResult.String())
	}
	submission.logger.LogTo(submission.currentGroup, "score", groupResult.String())
}

func (submission *Submission) checkSourceCodes() {
	err := ts.CheckSource(submission.Question, submission.solution, submission.Language)
	if err != nil {
		cobra.CheckErr(fmt.Errorf("check source code failed beucase %v", err))
	}

}

func (submission *Submission) RunSuite() *SubmissionResult {
	submission.PrintInitialInfo()

	submission.initSandboxWithoutTest()
	submission.checkReq()
	submission.compile()
	submission.backupCompiled()

	submission.checkSourceCodes()

	submission.Result = &SubmissionResult{
		Submission:       submission,
		TestGroupResults: submission.exploreTestGroups(),
	}

	for _, groupResult := range submission.Result.TestGroupResults {
		submission.RunGroup(groupResult)
	}

	submission.createZipResult()

	return submission.Result
}
