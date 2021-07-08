package run

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Submission struct {
	Language      string
	SourcePath    string
	SourceContent string
}

func NewSubmission(sourcePath string) (*Submission, error) {
	sourceContent, readErr := ioutil.ReadFile(sourcePath)
	if readErr != nil {
		return nil, readErr
	}

	submission := &Submission{
		SourcePath:    sourcePath,
		SourceContent: string(sourceContent),
		Language:      strings.Trim(filepath.Ext(sourcePath), "."),
	}
	return submission, nil
}
