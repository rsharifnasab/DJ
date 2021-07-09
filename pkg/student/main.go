package student

import (
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"

	"github.com/rsharifnasab/DJ/pkg/judge"
	"github.com/rsharifnasab/DJ/pkg/run"
	"github.com/rsharifnasab/DJ/pkg/util"
	"github.com/shirou/gopsutil/process"
)

func monitorMem(p *process.Process) {
	// there is also linux only solution with setrlimit:
	// read `man 2 prlimit`  and
	// https://golang.hotexamples.com/examples/syscall/-/Setrlimit/golang-setrlimit-function-examples.html
	// https://www.quora.com/Computer-Programming/What-is-the-simplest-and-most-accurate-way-to-measure-the-memory-used-by-a-program-in-a-programming-contest-environment/answer/Vivek-Prakash-2
	// the current solution should be platform independant
	// (why windows still exists?)

	// using https://pkg.go.dev/github.com/shirou/gopsutil/process#MemoryInfoStat

	for {
		// TODO: count all sub processes too
		memInfo, err := p.MemoryInfo()
		if err != nil {
			println("error eccoured")
			println(err.Error())
			return
		}
		mem := memInfo.RSS / 1024 / 1024
		println(mem)
		println(memInfo.String())
		if mem == 0 {
			println("zero mem, shuttong down")
			return
		} else if mem > 100000000 {
			err := p.Kill()
			if err != nil {
				panic(err)
			}
			println("killed by much mem usage")

		}

	}
}

func Run() {

	rules, err := judge.LoadRules("./examples/rules.yml")
	if err != nil {
		panic(err)
	}

	question, err := judge.NewQuestion("./examples/Q1", rules)
	if err != nil {
		panic(err)
	}
	util.PrintStruct(question)

	submission, submitErr := run.NewSubmission("./examples/solution.cpp")
	if submitErr != nil {
		panic(submitErr)
	}

	util.PrintStruct(submission)
	//println(submission.SourceContent)
	util.PrintStruct(question.AvailableLangs[submission.Language])
	println("\n----------\n")

	// which language
	// apply rules
	// compile, print compile errors
	// for: run on all test cases
	//   with tests count and out limit and ...

	const compiledFile = "./examples/a.out"

	const timeout = 2 * time.Second

	const testsCount = 7

	const outLimit = /* 1024 * 1024 * */ 10

	//for i := 1; i <= testsCount; i++ {
	for i := 2; i <= 2; i++ {
		// quera use this:
		// https://github.com/mjnaderi/Sharif-Judge/blob/Version-1/tester/runcode.sh
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
		stdoutPipe.Close()
		if err != nil {
			panic(err)
		}

		err = cmd.Start()
		if err != nil {
			panic(err)
		}

		pid := cmd.Process.Pid
		process, err := process.NewProcess(int32(pid))
		if err != nil {
			panic(err)
		}
		go monitorMem(process)

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
