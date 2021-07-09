package judge

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Rule struct {
	Description string   `yaml:"description"`
	Yes         []string `yaml:"yes"`
	No          []string `yaml:"no"`
	// TODO: add a sciprting language rule
}

const RulesRelativePath string = "rules.yml"

func (judge *Judge) LoadRules() error {
	rulesPath := judge.BasePath + RulesRelativePath

	yamlData, loadErr := ioutil.ReadFile(rulesPath)
	if loadErr != nil {
		return loadErr
	}

	judge.Rules = make(map[string]*Rule)

	unmarshalErr := yaml.UnmarshalStrict(yamlData, &judge.Rules)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	return nil
}

func (rule *Rule) Apply(source string) error {
	return nil
}
