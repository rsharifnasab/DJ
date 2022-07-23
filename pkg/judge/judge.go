package judge

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/plus3it/gorecurcopy"
	"github.com/rsharifnasab/DJ/pkg/run"
	"github.com/rsharifnasab/DJ/pkg/util"
	"github.com/spf13/cobra"
)

func checkReq(judgerPath string) {
	stdout, err := run.JustOut(judgerPath + "/req.sh")
	cobra.CheckErr(err)
	space := regexp.MustCompile(`\s+`)
	reqWords := space.Split(stdout, -1)
	run.CheckAndErrorRequirements(reqWords)
}

func runTestCase(judgerPath string, i int) (result TestRunResult) {
	result.Run = true

	command := fmt.Sprintf("%s/run.sh text %d", judgerPath, i)
	stdout, err := run.JustOut(command)
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
	//TODO: better handing
	switch stdout {
	case "pass":
		result.Pass = true
		return
	case "fail":
		return
	default:
		result.MalformedOutput = true
		return
	}
}

type TestRunResult struct {
	Run             bool
	Pass            bool
	Killed          bool
	TimedOut        bool
	NoResult        bool
	NonZero         bool
	MalformedOutput bool
}

func runAllTests(judgerPath string, n int) []TestRunResult {
	results := make([]TestRunResult, n)
	for i := 0; i < n; i++ {
		results[i] = runTestCase(judgerPath, i)
	}
	return results
}

func numberOfTests(judgerPath string) int {
	stdout, err := run.JustOut(judgerPath + "/run.sh count")
	cobra.CheckErr(err)
	n, err := strconv.Atoi(strings.TrimSpace(stdout))
	cobra.CheckErr(err)
	return n
}
func makeTempfolder() string {
	tmpFolder, err := ioutil.TempDir("", "djtmp*")
	cobra.CheckErr(err)
	return tmpFolder
}

func copyLibFiles(judgerPath, tmpFolder string) {
	err := gorecurcopy.CopyDirectory(judgerPath+"/lib", tmpFolder)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		cobra.CheckErr(err)
	}
}

func copyDir(srcPath, destPath string) {
	err := gorecurcopy.CopyDirectory(srcPath, destPath)
	cobra.CheckErr(err)
}

func compile(judgerPath string) {
	stdout, err := run.JustOut(judgerPath + "/run.sh compile")
	if err != nil {
		cobra.CheckErr(fmt.Errorf("cannot compile: %w", err))
	}
	println("compiler warnings: " + stdout)
}

func judge(submission run.Submission) (int, int) {
	tmpFolder := makeTempfolder()
	defer os.RemoveAll(tmpFolder)

	judgerPath := submission.Judger
	submissionPath := submission.Path
	questionPath := submission.Question
	_ = questionPath

	copyDir(judgerPath+"/lib", tmpFolder+"/lib")
	copyDir(judgerPath+"/project", tmpFolder)
	copyDir(submissionPath, tmpFolder)
	copyDir(judgerPath+"/test", tmpFolder)

	checkReq(judgerPath)
	compile(judgerPath)
	noOfTests := numberOfTests(judgerPath)
	results := runAllTests(judgerPath, noOfTests)
	util.PrintStruct(results)
	passed := 0
	for _, res := range results {
		if res.Pass {
			passed++
		}
	}
	return passed, noOfTests
}
