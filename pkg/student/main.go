package student

import (
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"

	"github.com/rsharifnasab/DJ/pkg/util"
	"gopkg.in/yaml.v2"
)

type TestCase struct {
	num int
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

	AllLangs       bool              `yaml:"all_langs"`
	AvailableLangs []*LanguageConfig `yaml:"languages"`
}

func GetDefaultQuestion(questionPath string) *Question {
	return &Question{
		Path: questionPath,

		MaxScore: 100,
		OutLimit: 10,

		AllLangs:       true,
		AvailableLangs: make([]*LanguageConfig, 0),

		Testcase: make([]*TestCase, 0),
	}
}

func NewQuestion(questionPath string, rules *Rules) (*Question, error) {
	yamlData, loadErr := ioutil.ReadFile(questionPath + "/config.yml")
	if loadErr != nil {
		return nil, fmt.Errorf("cannot load config.yml in %v because:\n\t %v",
			questionPath, loadErr)
	}
	question := GetDefaultQuestion(questionPath)

	unmarshalErr := yaml.UnmarshalStrict(yamlData, &question)
	if unmarshalErr != nil {
		return nil, fmt.Errorf("cannot unmarshal config.yml file because:\n\t %v",
			unmarshalErr)
	}

	//util.PrintStruct(question.AvailableLangs)
	for _, lang := range question.AvailableLangs {
		lang.Rules = make(map[string]*Rule, 0)
		for _, ruleName := range lang.RuleNames {
			ruleObj, foundRule := rules.Map[ruleName]
			if !foundRule {
				return nil, fmt.Errorf("rulename [%v] doesn't exist", ruleName)
			}
			lang.Rules[ruleName] = &ruleObj
		}
	}
	return question, nil
}

type Rule struct {
	Description string   `yaml:"description"`
	Yes         []string `yaml:"yes"`
	No          []string `yaml:"no"`
}

type Rules struct {
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
	util.PrintStruct(rules)

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
