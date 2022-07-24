package judge

type Submission struct {
	Solution      string
	Judger        string
	Question      string
	sandboxDir    string
	CompiledState string
}

func NewSubmission(path, runner, question string) *Submission {
	submission := &Submission{
		Solution: path,
		Judger:   runner,
		Question: question,
	}
	return submission
}

type TestResult struct {
	Run             bool
	Pass            bool
	Killed          bool
	TimedOut        bool
	NoResult        bool
	NonZero         bool
	MalformedOutput bool
}

type TestGroupResult struct {
	TestCount   int
	TestResults []*TestResult
	Name        string
}

type SubmissionResult struct {
	Submission       *Submission
	TestGroupResults []*TestGroupResult
}
