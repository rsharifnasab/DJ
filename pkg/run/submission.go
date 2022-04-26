package run

type Submission struct {
	Path     string
	Runner   string
	Question string
}

func NewSubmission(path, runner, question string) *Submission {

	submission := &Submission{
		Path:     path,
		Runner:   runner,
		Question: question,
	}
	return submission
}
