package run

type Submission struct {
	Path     string
	Judger   string
	Question string
}

func NewSubmission(path, runner, question string) *Submission {
	submission := &Submission{
		Path:     path,
		Judger:   runner,
		Question: question,
	}
	return submission
}
