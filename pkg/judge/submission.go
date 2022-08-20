package judge

import (
	"fmt"
	"math"
	"strings"

	"github.com/rsharifnasab/DJ/pkg/util"
)

type Submission struct {
	Solution      string
	Judger        string
	Question      string
	Language      string
	sandboxDir    string
	CompiledState string
	Result        string
	currentGroup  string
	logger        *util.Logger
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
func (r *TestResult) String() string {
	if !r.Run {
		return "not ran yet"
	} else if r.Pass {
		return "passed"
	} else if r.Killed {
		return "killed"
	} else if r.TimedOut {
		return "timed out"
	} else if r.NoResult {
		return "printed nothing"
	} else if r.NonZero {
		return "exited with exit-code != 0"
	} else if r.MalformedOutput {
		return "malformed output"
	} else if r.Wrong {
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
	builder.WriteString("[")
	for i, e := range gr.TestResults {
		str := fmt.Sprintf("   test %d: %s\n", i+1, e.String())
		builder.WriteString(str)
	}

	str := fmt.Sprintf("] testgroup [%s]: (%d/%d) - %d%%", gr.Name, gr.PassedCount(), gr.AllCount(), gr.Score())
	builder.WriteString(str)

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
