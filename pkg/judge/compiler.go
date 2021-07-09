package judge

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Compiler struct {
	Name                   string `yaml:"name"`
	TemplateCompileCommand string `yaml:"compile"`
	TemplateRunCommand     string `yaml:"run"`
}

const CompilerConfigRelativePath string = "compiler.yml"

func (judge *Judge) LoadCompilers() error {
	compilersPath := judge.BasePath + CompilerConfigRelativePath

	yamlData, loadErr := ioutil.ReadFile(compilersPath)
	if loadErr != nil {
		return loadErr
	}

	judge.Compilers = make([]*Compiler, 0)

	unmarshalErr := yaml.UnmarshalStrict(yamlData, &judge.Compilers)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	// todo check all compilers installed

	return nil
}
