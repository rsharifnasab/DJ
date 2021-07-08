package judge

import "time"

type LanguageConfig struct {
	TimeLimit   time.Duration    `yaml:"time"`
	MemoryLimit int              `yaml:"memory"`
	RuleNames   []string         `yaml:"rules"`
	Rules       map[string]*Rule `yaml:"NONE"`
}
