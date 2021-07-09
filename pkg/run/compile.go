package run

import (
	"github.com/rsharifnasab/DJ/pkg/judge"
)

type Language struct {
	Name                   string `yaml:"name"`
	TemplateCompileCommand string `yaml:"compile"`
	TemplateRunCommand     string `yaml:"run"`
}

type Judge struct {
	Languages []*Language
	Rules     map[string]*judge.Rule
}
