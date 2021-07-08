package judge

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Question struct {
	Name string `yaml:"name"`
	Path string

	MaxScore int `yaml:"max_score"`
	OutLimit int `yaml:"out_limit"`

	Testcase []*TestCase

	AvailableLangs []*LanguageConfig `yaml:"languages"`
}

func GetDefaultQuestion(questionPath string) *Question {
	return &Question{
		Path: questionPath + "/",

		MaxScore: 100,
		OutLimit: 10,

		AvailableLangs: make([]*LanguageConfig, 0),

		Testcase: make([]*TestCase, 0),
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

func NewQuestion(questionPath string, rules *Rules) (*Question, error) {
	question := GetDefaultQuestion(questionPath)

	if loadErr := question.ReadConfigFile(); loadErr != nil {
		return nil, loadErr
	}

	if convertErr := question.ruleNameToRule(rules); convertErr != nil {
		return nil, convertErr
	}

	if testCaseErr := question.loadTestCases(); testCaseErr != nil {
		return nil, testCaseErr
	}

	return question, nil
}
