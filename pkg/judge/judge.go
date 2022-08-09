package judge

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/rsharifnasab/DJ/pkg/run"
	"github.com/rsharifnasab/DJ/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func checkReq(submission *Submission) {
	stdout, err := run.JustOut(submission.sandboxDir + "/req.sh")
	cobra.CheckErr(err)
	space := regexp.MustCompile(`\s+`)
	reqWords := space.Split(stdout, -1)
	run.CheckAndErrorRequirements(reqWords)
}

func runTestCase(submission *Submission, i int) (result *TestResult) {

	command := fmt.Sprintf("%s/run.sh test %d", submission.sandboxDir, i)
	stdout, stderr, err := run.DefaultRun(command)

	// TODO: write stderr to file
	util.LogToResult(submission.Result, submission.currentGroup, strconv.Itoa(i), stderr)

	result = &TestResult{
		Run: true,
	}
	if err != nil {
		switch err {
		case run.MemoryLimitError:
			result.Killed = true
			return
		case run.OutputLimitError:
			result.Killed = true
			return
		case run.MalformedCommandError:
			panic(err)
		case run.TimedOutError:
			result.TimedOut = true
			return
		case run.NoOutputError:
			result.NoResult = true
			return
		case run.NotValidExecutableError:
			panic("cannot run test " + strconv.Itoa(i))
		case run.NonZeroExitError:
			result.NonZero = true
			return
		}
	} else {
		//fmt.Println("no error")
	}
	//TODO: better handling
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

func testCount(submission *Submission) int {
	command := fmt.Sprintf("%s/run.sh count", submission.sandboxDir)
	stdout, err := run.JustOut(command)
	cobra.CheckErr(err)
	n, err := strconv.Atoi(strings.TrimSpace(stdout))
	cobra.CheckErr(err)
	return n
}

func compile(submission *Submission) {
	command := fmt.Sprintf("%s/run.sh compile", submission.sandboxDir)

	stdout, stderr, err := run.DefaultRun(command)
	if err != nil && err != run.NoOutputError {
		cobra.CheckErr(fmt.Errorf("Compilation failed:\nerr: %w", err))
	} else {
		fmt.Println("Compilation successful")
	}
	util.LogToResult(submission.Result, "", "compile", stdout)
	util.LogToResult(submission.Result, "", "compile", stderr)
}

func initFolderWithoutTest(submission *Submission) {
	util.CopyDir(submission.Judger, submission.sandboxDir)
	util.CopyDir(submission.Solution, submission.sandboxDir+"/src")
}

func backupCompiled(submission *Submission) {
	util.CopyDir(submission.sandboxDir, submission.CompiledState)
}

func restoreCompiled(submission *Submission) {
	err := os.RemoveAll(submission.sandboxDir)
	cobra.CheckErr(err)
	util.CopyDir(submission.CompiledState, submission.sandboxDir)
}

func exploreTestGroups(submission *Submission) []*TestGroupResult {
	res := make([]*TestGroupResult, 0, 10)

	files, err := ioutil.ReadDir(submission.Question)
	cobra.CheckErr(err)
	for _, f := range files {
		testGroup := &TestGroupResult{
			Name:        f.Name(), // just name
			TestResults: make([]*TestResult, 0, 10),
		}
		res = append(res, testGroup)
	}
	return res
}

func prepareTestGroup(submission *Submission, groupName string) {
	src := fmt.Sprintf("%s/%s", submission.Question, groupName)
	dest := fmt.Sprintf("%s/testgroup", submission.sandboxDir)
	util.CopyDir(src, dest)
}

func RunSubmission(submission *Submission) *SubmissionResult {
	submission.sandboxDir = util.MakeTempfolder()
	submission.CompiledState = util.MakeTempfolder()
	if !viper.GetBool("debug") {
		defer os.RemoveAll(submission.sandboxDir)
		defer os.RemoveAll(submission.CompiledState)
	}

	if submission.Result == "" {
		// reminder: we don't want to remove the result folder!
		submission.Result = util.MakeTempfolder()
	}
	fmt.Printf("result dir: %v\n", submission.Result)

	if viper.GetBool("debug") {
		fmt.Printf("Sandbox  dir: %s\n", submission.sandboxDir)
		fmt.Printf("compiled dir: %s\n", submission.CompiledState)
	}

	initFolderWithoutTest(submission)
	checkReq(submission)

	compile(submission)
	backupCompiled(submission)

	submResult := &SubmissionResult{
		Submission:       submission,
		TestGroupResults: exploreTestGroups(submission),
	}

	for _, groupResult := range submResult.TestGroupResults {
		restoreCompiled(submission)
		prepareTestGroup(submission, groupResult.Name)
		submission.currentGroup = groupResult.Name // for logging
		groupResult.TestCount = testCount(submission)
		for i := 1; i <= groupResult.TestCount; i++ {
			testResult := runTestCase(submission, i)
			groupResult.TestResults = append(groupResult.TestResults, testResult)
		}
		fmt.Println(groupResult.String())
		//time.Sleep(5 * time.Second)

	}

	return submResult

}
