package question

import (
	"fmt"
)

const (
	testsFolder = "tests/"
	testsIn     = testsFolder + "in/"
	testsOut    = testsFolder + "out/"
	configFile  = "config.yml"
)

func (question *Question) GetTestsFolder() string {
	return question.Path + testsFolder
}

func (question *Question) GetTestsInputFolder() string {
	return question.Path + testsIn
}

func (question *Question) GetTestsOutputFolder() string {
	return question.Path + testsOut
}

func (question *Question) GetInputFilePath(num int) string {
	return fmt.Sprintf("%vinput%d.txt", question.GetTestsInputFolder(), num)
}

func (question *Question) GetOutputFilePath(num int) string {
	return fmt.Sprintf("%voutput%d.txt", question.GetTestsOutputFolder(), num)
}

func GetTestNumFromInputFileName(name string) (int, error) {
	var num int
	n, scanfErr := fmt.Sscanf(name, "input%d.txt", &num)
	if n != 1 || scanfErr != nil {
		return 0, fmt.Errorf("%v : filename malformed", name)
	} else {
		return num, nil
	}
}
