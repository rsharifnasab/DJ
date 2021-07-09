package question

import (
	"fmt"
	"io/ioutil"

	"github.com/rsharifnasab/DJ/pkg/judge"
	"gopkg.in/yaml.v2"
)

type Question struct {
	Name string `yaml:"name"`
	Path string

	MaxScore int `yaml:"max_score"`
	OutLimit int `yaml:"out_limit"`

	Testcase []*TestCase

	AvailableLangs map[string]*LanguageConfig `yaml:"languages"`
}

func GetDefaultQuestion(questionPath string) *Question {
	return &Question{
		Path: questionPath + "/",

		MaxScore: 100,
		OutLimit: 10,

		Testcase: make([]*TestCase, 0),

		AvailableLangs: make(map[string]*LanguageConfig, 0),
	}
}

func (question *Question) ReadConfigFile() error {
	effectiveConfigPath := question.Path + configFile

	yamlData, loadErr := ioutil.ReadFile(effectiveConfigPath)
	if loadErr != nil {
		return loadErr
	}
	unmarshalErr := yaml.UnmarshalStrict(yamlData, &question)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	return nil
}

func (question *Question) lookuplanguageconfigs(judge *judge.Judge) error {
	return nil
}

func NewQuestion(questionPath string, judge *judge.Judge) (*Question, error) {
	question := GetDefaultQuestion(questionPath)

	if loadErr := question.ReadConfigFile(); loadErr != nil {
		return nil, loadErr
	}

	if convertErr := question.ruleNameToRule(judge.Rules); convertErr != nil {
		return nil, convertErr
	}

	if testCaseErr := question.loadTestCases(); testCaseErr != nil {
		return nil, testCaseErr
	}

	if langErr := question.lookuplanguageconfigs(judge); langErr != nil {
		return nil, langErr
	}

	return question, nil
}

func (question *Question) ruleNameToRule(rules map[string]*judge.Rule) error {
	for _, lang := range question.AvailableLangs {
		lang.Rules = make(map[string]*judge.Rule, 0)
		for _, ruleName := range lang.RuleNames {
			ruleObj, foundRule := rules[ruleName]
			if !foundRule {
				return fmt.Errorf("rulename %v doesn't exist", ruleName)
			}
			lang.Rules[ruleName] = ruleObj
		}
	}
	return nil
}
