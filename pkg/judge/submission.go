package judge

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/rsharifnasab/DJ/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Submission struct {
	Judge        string
	Question     string
	Language     string
	UserSolution string
	ResultDir    string

	solution      string
	sandboxDir    string
	CompiledState string
	currentGroup  string
	logger        *util.Logger

	Result *SubmissionResult
}

func (submission *Submission) initFields() {
	if ok, err := util.IsZip(submission.UserSolution); err != nil {
		cobra.CheckErr(err)
	} else if ok {
		submission.solution = util.MakeTempfolder()
		util.Unzip(submission.UserSolution, submission.solution)
		if viper.GetBool("debug") {
			fmt.Printf("unzipped to submission.Solution: %s\n", submission.solution)
		}
	} else {
		submission.solution = submission.UserSolution
	}
	submission.solution = util.AutoCd(submission.solution)

	if submission.Language == "" {
		if lang, err := util.AutoDetectLanguage(submission.solution); err != nil {
			submission.Language = "generic"
		} else {
			submission.Language = lang
		}
	}

	submission.sandboxDir = util.MakeTempfolder()
	submission.CompiledState = util.MakeTempfolder()
}

func (submission *Submission) initLogger() {
	if submission.ResultDir == "" {
		// reminder: we don't want to remove the result folder!
		submission.ResultDir = util.MakeTempfolder()
	}
	submission.logger = util.NewLogger(submission.ResultDir)
	submission.ResultDir = submission.logger.BasePath
}

func NewSubmission(judge, question, language, userSolution, resultDir string) *Submission {
	s := &Submission{
		Judge:        judge,
		Question:     question,
		Language:     language,
		UserSolution: userSolution,
		ResultDir:    resultDir,
	}
	s.initFields()
	s.initLogger()

	return s
}

type TestResult struct {
	Run             bool
	Pass            bool
	Wrong           bool
	Killed          bool
	TimedOut        bool
	NoResult        bool
	NonZero         bool
	MalformedOutput bool
}

func (testResult *TestResult) isPassed() bool {
	return testResult.Pass
}
func (testResult *TestResult) String() string {
	if !testResult.Run {
		return "not ran yet"
	} else if testResult.Pass {
		return "passed"
	} else if testResult.Killed {
		return "killed"
	} else if testResult.TimedOut {
		return "timed out"
	} else if testResult.NoResult {
		return "printed nothing"
	} else if testResult.NonZero {
		return "exited with exit-code != 0"
	} else if testResult.MalformedOutput {
		return "malformed output"
	} else if testResult.Wrong {
		return "wrong answer"
	} else {
		panic("malformed testResult")
	}
}

type TestGroupResult struct {
	Name        string
	TestCount   int
	TestResults []*TestResult
}

func (gr *TestGroupResult) PassedCount() int {
	counter := 0
	for _, e := range gr.TestResults {
		if e.isPassed() {
			counter++
		}
	}
	return counter
}

func (gr *TestGroupResult) AllCount() int {
	return gr.TestCount
}

func (gr *TestGroupResult) Score() int {
	if gr.AllCount() == 0 {
		return 0
	}
	return int(math.Ceil(100 * float64(gr.PassedCount()) / float64(gr.AllCount())))
}

func (gr *TestGroupResult) String() string {
	builder := strings.Builder{}
	builder.WriteString(
		fmt.Sprintf("testgroup [%s]: (%d/%d) - %d%%  [\n",
			gr.Name,
			gr.PassedCount(),
			gr.AllCount(),
			gr.Score(),
		),
	)
	for i, e := range gr.TestResults {
		str := fmt.Sprintf("   test %d: %s\n", i+1, e.String())
		builder.WriteString(str)
	}

	builder.WriteString("]")

	return builder.String()
}

type SubmissionResult struct {
	Submission       *Submission
	TestGroupResults []*TestGroupResult
}

func (sr *SubmissionResult) String() string {
	builder := strings.Builder{}
	builder.WriteString("Submission Result: [\n")
	for _, gr := range sr.TestGroupResults {
		str := fmt.Sprintf("    testgroup %s: (%d/%d) - %d%%\n", gr.Name, gr.PassedCount(), gr.AllCount(), gr.Score())
		builder.WriteString(str)
	}
	builder.WriteString("]")
	return builder.String()
}

func (sr *SubmissionResult) DumpTo(file string) error {
	str := sr.String()
	err := os.WriteFile(file, []byte(str), 0666)
	return err
}
