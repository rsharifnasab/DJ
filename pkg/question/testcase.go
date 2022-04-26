package question

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type TestCase struct {
	Num int
	// TODO: weight

	InputFile  string
	OutputFile string
}

func (question *Question) NewTestCase(num int) *TestCase {
	return &TestCase{
		Num:        num,
		InputFile:  question.GetInputFilePath(num),
		OutputFile: question.GetOutputFilePath(num),
	}
}

func (question *Question) CreateTestcases(count int) {
	for i := 1; i <= count; i++ {
		question.Testcase = append(question.Testcase, question.NewTestCase(i))
	}

}

func CheckTestCasesInOrder(tests []int) error {
	sort.Ints(tests)
	if first := tests[0]; first != 1 {
		return fmt.Errorf("first tests isn't 1, it's %v", first)
	} else if last := tests[len(tests)-1]; last != len(tests) {
		return fmt.Errorf("tests are not continues, last test is: %v", last)
	} else {
		return nil
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func (question *Question) loadTestCases() error {
	inputFiles, readDirErr := ioutil.ReadDir(question.GetTestsInputFolder())
	if readDirErr != nil {
		return readDirErr
	}

	tests := make([]int, 0, 20)

	for _, v := range inputFiles {
		inpName := v.Name()

		if !exists(question.GetTestsInputFolder() + inpName) {
			return fmt.Errorf("input file %v doesn't exist", inpName)
		}

		num, err := GetTestNumFromInputFileName(inpName)
		if err != nil {
			return err
		}
		if !exists(question.GetOutputFilePath(num)) {
			return fmt.Errorf("output file %v doesn't exist", inpName)
		}

		tests = append(tests, num)
	}

	orderErr := CheckTestCasesInOrder(tests)
	if orderErr != nil {
		return orderErr
	}

	// everything OK!
	question.CreateTestcases(len(tests))

	return nil
}
