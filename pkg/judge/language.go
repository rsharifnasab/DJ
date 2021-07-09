package judge

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Language struct {
	Name                   string `yaml:"name"`
	TemplateCompileCommand string `yaml:"compile"`
	TemplateRunCommand     string `yaml:"run"`
}

const LanguagesRelativePath string = "languages.yml"

func (judge *Judge) LoadLanguages() error {
	languagesPath := judge.BasePath + LanguagesRelativePath

	yamlData, loadErr := ioutil.ReadFile(languagesPath)
	if loadErr != nil {
		return loadErr
	}

	judge.Languages = make([]*Language, 0)

	unmarshalErr := yaml.UnmarshalStrict(yamlData, &judge.Languages)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	return nil
}
