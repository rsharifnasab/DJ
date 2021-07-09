package run

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Submission struct {
	LanguageName string
	//Language     *judge.Language

	SourcePath    string
	SourceContent string
}

func NewSubmission(sourcePath string) (*Submission, error) {
	// TODO: get question and select proper LanguageConfig from file type
	sourceContent, readErr := ioutil.ReadFile(sourcePath)
	if readErr != nil {
		return nil, readErr
	}

	submission := &Submission{
		SourcePath:    sourcePath,
		SourceContent: string(sourceContent),
		LanguageName:  strings.Trim(filepath.Ext(sourcePath), "."),
	}
	return submission, nil
}
