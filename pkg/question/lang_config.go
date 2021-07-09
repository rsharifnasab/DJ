package question

import (
	"time"

	"github.com/rsharifnasab/DJ/pkg/judge"
)

type LanguageConfig struct {
	TimeLimit   time.Duration `yaml:"time"`
	MemoryLimit int           `yaml:"memory"`

	RuleNames []string               `yaml:"rules"`
	Rules     map[string]*judge.Rule `yaml:"NONE"`

	Compiler *judge.Compiler `yaml:"NONE2"`
}
