package judge

import (
	"fmt"
	"io/ioutil"

	"github.com/rsharifnasab/DJ/pkg/run"
	"gopkg.in/yaml.v2"
)

type Rule struct {
	Description string   `yaml:"description"`
	Yes         []string `yaml:"yes"`
	No          []string `yaml:"no"`
	// TODO: add a sciprting language rule
}

func LoadRules(rulesPath string) (*run.Rules, error) {
	yamlData, loadErr := ioutil.ReadFile(rulesPath)

	if loadErr != nil {
		return nil, loadErr
	}
	rules := make(map[string]*Rule)

	unmarshalErr := yaml.UnmarshalStrict(yamlData, &rules)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return rules, nil
}

func (question *Question) ruleNameToRule(rules *Rules) error {
	for _, lang := range question.AvailableLangs {
		lang.Rules = make(map[string]*Rule, 0)
		for _, ruleName := range lang.RuleNames {
			ruleObj, foundRule := rules.Map[ruleName]
			if !foundRule {
				return fmt.Errorf("rulename %v doesn't exist", ruleName)
			}
			lang.Rules[ruleName] = &ruleObj
		}
	}
	return nil
}

func (rule *Rule) Apply(source string) error {

}
