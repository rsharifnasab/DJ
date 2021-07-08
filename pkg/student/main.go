package student

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"

	"github.com/rsharifnasab/DJ/pkg/util"
	"gopkg.in/yaml.v2"
)

type TestCase struct {
	Num int
	// TODO: weight

	InputFile  string
	OutputFile string
}

type LanguageConfig struct {
	Name        string           `yaml:"lang"`
	TimeLimit   time.Duration    `yaml:"time"`
	MemoryLimit int              `yaml:"memory"`
	RuleNames   []string         `yaml:"rules"`
	Rules       map[string]*Rule `yaml:"NONE"`
}

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
		Path: questionPath,

		MaxScore: 100,
		OutLimit: 10,

		AvailableLangs: make([]*LanguageConfig, 0),

		Testcase: make([]*TestCase, 0),
	}
}

func (question *Question) ruleNameToRule(rules *Rules) error {
	for _, lang := range question.AvailableLangs {
		lang.Rules = make(map[string]*Rule, 0)
		for _, ruleName := range lang.RuleNames {
			ruleObj, foundRule := rules.Map[ruleName]
			if !foundRule {
				return fmt.Errorf("rulename [%v] doesn't exist", ruleName)
			}
			lang.Rules[ruleName] = &ruleObj
		}
	}
	return nil
}

const (
	testsFolder = "tests/"
	testsIn     = testsFolder + "in/"
	testsOut    = testsFolder + "out/"
)

func (question *Question) GetTestsFolder() string {
	return question.Path + "/" + testsFolder
}

func (question *Question) GetTestsInputFolder() string {
	return question.Path + "/" + testsIn
}

func (question *Question) GetTestsOutputFolder() string {
	return question.Path + "/" + testsOut
}

func (question *Question) GetInputFilePath(num int) string {
	return fmt.Sprintf("%vinput%d.txt", question.GetTestsInputFolder(), num)
}

func (question *Question) GetOutputFilePath(num int) string {
	return fmt.Sprintf("%voutput%d.txt", question.GetTestsOutputFolder(), num)
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

func GetTestNumFromInputFileName(name string) (int, error) {
	var num int
	n, scanfErr := fmt.Sscanf(name, "input%d.txt", &num)
	if n != 1 || scanfErr != nil {
		return 0, fmt.Errorf("%v : filename malformed", name)
	} else {
		return num, nil
	}
}

func CheckFileExists(path string) error {
	if stat, err := os.Stat(path); err != nil {
		return fmt.Errorf("problem with file %v : %v", path, err)
	} else if stat.IsDir() {
		return fmt.Errorf("%v is a directory", path)
	} else {
		return nil
	}
}

func CheckTestCasesInOrder(tests []int) error {
	sort.Ints(tests)
	if first := tests[0]; first != 1 {
		return fmt.Errorf("first tests isn't 1, bit it's %v", first)
	} else if last := tests[len(tests)-1]; last != len(tests) {
		return fmt.Errorf("tests are not continues, last test is: %v", last)
	} else {
		return nil
	}
}

func (question *Question) loadTestCases() error {
	inputFiles, readDirErr := ioutil.ReadDir(question.GetTestsInputFolder())
	if readDirErr != nil {
		return readDirErr
	}

	tests := make([]int, 0, 20)

	for _, v := range inputFiles {
		inpName := v.Name()

		inputFileErr := CheckFileExists(question.GetTestsInputFolder() + inpName)
		if inputFileErr != nil {
			return inputFileErr
		}

		num, err := GetTestNumFromInputFileName(inpName)
		if err != nil {
			return err
		}

		outputFileErr := CheckFileExists(question.GetOutputFilePath(num))
		if outputFileErr != nil {
			return outputFileErr
		}

		tests = append(tests, num)
	}

	orderErr := CheckTestCasesInOrder(tests)
	if orderErr != nil {
		return orderErr
	}
	question.CreateTestcases(len(tests))
	return nil
}

func NewQuestion(questionPath string, rules *Rules) (*Question, error) {
	question := GetDefaultQuestion(questionPath)

	yamlData, loadErr := ioutil.ReadFile(questionPath + "/config.yml")
	if loadErr != nil {
		return nil, fmt.Errorf("cannot load config.yml in %v because:\n\t %v",
			questionPath, loadErr)
	}
	unmarshalErr := yaml.UnmarshalStrict(yamlData, &question)
	if unmarshalErr != nil {
		return nil, fmt.Errorf("cannot unmarshal config.yml file because:\n\t %v",
			unmarshalErr)
	}

	if convertErr := question.ruleNameToRule(rules); convertErr != nil {
		return nil, convertErr
	}

	if testCaseErr := question.loadTestCases(); testCaseErr != nil {
		return nil, testCaseErr
	}

	return question, nil
}

type Rule struct {
	Description string   `yaml:"description"`
	Yes         []string `yaml:"yes"`
	No          []string `yaml:"no"`
}

type Rules struct {
	// TODO: apply rules in smart way with chroma
	// https://github.com/alecthomas/chroma#supported-languages
	Map map[string]Rule `yaml:"rules"`
}

func LoadRules(rulesPath string) (*Rules, error) {
	yamlData, loadErr := ioutil.ReadFile(rulesPath)

	if loadErr != nil {
		return nil, fmt.Errorf("cannot load %v because:\n\t %v",
			rulesPath, loadErr)
	}
	rules := &Rules{}

	unmarshalErr := yaml.UnmarshalStrict(yamlData, &rules)
	if unmarshalErr != nil {
		return nil, fmt.Errorf("cannot unmarshal rules yml file because:\n\t %v",
			unmarshalErr)
	}

	return rules, nil
}

func Run() {

	rules, err := LoadRules("./examples/rules.yml")
	if err != nil {
		panic(err)
	}
	//util.PrintStruct(rules)

	question, err := NewQuestion("./examples/Q1", rules)
	if err != nil {
		panic(err)
	}
	util.PrintStruct(question)

	const submissionFile = "./examples/solution.cpp"
	_ = submissionFile
	// todo compile

	const compiledFile = "./examples/a.out"

	const timeout = 2 * time.Second

	const testsCount = 7

	const outLimit = /* 1024 * 1024 * */ 10

	for i := 1; i <= testsCount; i++ {
		testInpAddr := fmt.Sprintf("./examples/Q1/tests/in/input%d.txt", i)
		testInpData, err := ioutil.ReadFile(testInpAddr)
		if err != nil {
			panic(err)
		}
		fmt.Printf("test input : [%v]\n", strings.TrimSpace(string(testInpData)))

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel() // cleanup resources

		// Create the command with our context
		cmd := exec.CommandContext(ctx, compiledFile)

		stdinWriter, err := cmd.StdinPipe()
		stdinWriter.Write(testInpData)
		if err != nil {
			panic(err)
		}

		stdoutPipe, err := cmd.StdoutPipe()
		if err != nil {
			panic(err)
		}

		err = cmd.Start()
		if err != nil {
			panic(err)
		}

		outBuf := make([]byte, outLimit+1)
		n, err := stdoutPipe.Read(outBuf)

		if n == outLimit+1 {
			println("output limit exceed")
		}
		out := outBuf[:n]

		// finished flag become true
		// and get probable error
		err = cmd.Wait()

		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("time limit exceed")
		} else {
			outStr := strings.TrimSpace(string(out))
			fmt.Printf("Output: [%v]\n", outStr)
		}

		if err != nil {
			fmt.Printf("abnormal status : [%v]\n", err)
		} else {
			fmt.Println("normal exit")
		}

		println()
	}
}

/*

	direct write to stdin of program
	if _, err = io.WriteString(stdin, "0\n"); err != nil {
		panic(err)
	}


	pipe a file to stdin

	rawFile, err := os.Open(testAddr)
	if err != nil {
		panic(err)
	}

	bufReader := bufio.NewReader(rawFile)


	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	stdinBuf := bufio.NewWriter(stdin)
	stdinBuf.ReadFrom(bufReader)

*/
