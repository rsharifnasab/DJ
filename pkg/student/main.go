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

type Preprocess struct {
}

type LanguageConfig struct {
	Name        string        `yaml:"lang"`
	TimeLimit   time.Duration `yaml:"time"`
	MemoryLimit int           `yaml:"memory"`
	Preprocess  Preprocess
}

type Question struct {
	Name string `yaml:"name"`
	Path string

	MaxScore int `yaml:"max_score"`
	OutLimit int `yaml:"out_limit"`

	Testcase []TestCase

	AvailableLangs []LanguageConfig `yaml:"languages"`
}

func NewQuestion(questionPath string) (*Question, error) {
	yamlData, loadErr := ioutil.ReadFile(questionPath + "/config.yml")
	if loadErr != nil {
		return nil, fmt.Errorf("cannot load config.yml in %v because %v",
			questionPath, loadErr)
	}
	question := &Question{
		Path: questionPath,
	}
	unmarshalErr := yaml.UnmarshalStrict(yamlData, &question)
	if unmarshalErr != nil {
		return nil, fmt.Errorf("cannot unmarshal yml file because: %v",
			unmarshalErr)
	}

	return question, nil
}

func Run() {
	question, err := NewQuestion("./examples/Q1")
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
		testInpAddr := fmt.Sprintf("./examples/tests/in/input%d.txt", i)
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
