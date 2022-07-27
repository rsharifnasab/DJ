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
	stdout, err := run.JustOut(command)
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
	fmt.Println("sandbox dir : " + submission.sandboxDir)
	command := fmt.Sprintf("%s/run.sh compile", submission.sandboxDir)

	err := run.JustRun(command)
	if err != nil && err != run.NoOutputError {
		cobra.CheckErr(fmt.Errorf("Compilation failed:\nerr: %w", err))
	}
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
	//fmt.Println("copy done")
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
	defer os.RemoveAll(submission.sandboxDir)
	submission.CompiledState = util.MakeTempfolder()
	defer os.RemoveAll(submission.CompiledState)

	// for each testgroup: testgroup to /testgroup
	// TODO
	initFolderWithoutTest(submission)
	checkReq(submission)

	compile(submission)
	fmt.Println("compile successful")
	backupCompiled(submission)

	submResult := &SubmissionResult{
		Submission:       submission,
		TestGroupResults: exploreTestGroups(submission),
	}

	for _, groupResult := range submResult.TestGroupResults {
		//fmt.Println(" - - - - - - - testgroup: " + groupResult.Name + " - - - - - - - - -")
		restoreCompiled(submission)
		//fmt.Println("restore compile success")
		prepareTestGroup(submission, groupResult.Name)
		//fmt.Println("prepare test group success")
		groupResult.TestCount = testCount(submission)
		for i := 1; i <= groupResult.TestCount; i++ {
			//fmt.Println(" - - - testgroup[" + groupResult.Name + "] test" + strconv.Itoa(i) + " - - -")
			testResult := runTestCase(submission, i)
			//util.PrintStruct(testResult)
			groupResult.TestResults = append(groupResult.TestResults, testResult)
		}
		fmt.Println(groupResult.String())
		//time.Sleep(5 * time.Second)

	}

	return submResult

}
