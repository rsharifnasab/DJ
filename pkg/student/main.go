package student

import (
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"

	"github.com/rsharifnasab/DJ/pkg/judge"
	"github.com/rsharifnasab/DJ/pkg/question"
	"github.com/rsharifnasab/DJ/pkg/run"
	"github.com/rsharifnasab/DJ/pkg/util"
	"github.com/shirou/gopsutil/process"
)

// recursive function to calculate whole process+childs memory usage
func TotalMemoryUsage(p *process.Process) (uint64, error) {
	currentMem, err := p.MemoryInfo()
	if err != nil {
		return 0, err
	}

	childs, err := p.Children()
	if err != nil { // doesn't have child
		return currentMem.RSS, nil
	}

	sum := currentMem.RSS
	for _, child := range childs {
		childUsage, err := TotalMemoryUsage(child)
		if err != nil {
			return sum, err
		}
		sum += childUsage
	}

	return sum, nil
}

func monitorMem(p *process.Process, memLimit uint64) {
	// there is also linux only solution with setrlimit:
	// read `man 2 prlimit`  and
	// https://golang.hotexamples.com/examples/syscall/-/Setrlimit/golang-setrlimit-function-examples.html
	// https://www.quora.com/Computer-Programming/What-is-the-simplest-and-most-accurate-way-to-measure-the-memory-used-by-a-program-in-a-programming-contest-environment/answer/Vivek-Prakash-2
	// the current solution should be platform independant
	// (why windows still exists?)

	// using https://pkg.go.dev/github.com/shirou/gopsutil/process#MemoryInfoStat

	for {
		totalUsingMem, err := TotalMemoryUsage(p)
		if err != nil {
			if _, ok := err.(*fs.PathError); ok {
				//fmt.Printf("err type : %T\n err val : %v\nerror text : %v\n", fsErr, fsErr, fsErr.Error())
				return
			} else {
				panic(err)
			}
		}
		println(totalUsingMem)
		if totalUsingMem > memLimit {
			err := p.Kill()
			if err != nil {
				panic(err)
			}
			println("killed by much mem usage")

		}
		time.Sleep(100 * time.Millisecond)
	}
}

func Run() {

	judge, err := judge.InitJudge("./examples")
	if err != nil {
		panic(err)
	}
	util.PrintStruct(judge)

	question, err := question.NewQuestion("./examples/Q1", judge)
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
	util.PrintStruct(question.AvailableLangs[submission.LanguageName])
	println("\n----------\n")

	//runExampleTests()
}

func runExampleTests() {
	// which language
	// apply rules
	// compile, print compile errors
	// for: run on all test cases
	//   with tests count and out limit and ...

	const compiledFile = "./examples/a.out"

	const timeout = 2 * time.Second

	const testsCount = 7

	const outLimit = /* 1024 * 1024 * */ 10

	const memLimit = 12 * 1024 * 1024

	//for i := 1; i <= testsCount; i++ {
	for i := 1; i <= 7; i++ {
		// sharif-judge use this:
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
		go monitorMem(process, memLimit)

		outBuf := make([]byte, outLimit+1)
		n, err := stdoutPipe.Read(outBuf)

		if n == outLimit+1 {
			println("output limit exceed")
		}
		out := outBuf[:n]

		err = stdoutPipe.Close()
		if err != nil {
			panic(err)
		}
		// TODO check how it is working
		// and how HEAD works

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
