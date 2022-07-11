package run

import (
	"context"
	"io/fs"
	"os/exec"
	"syscall"
	"time"

	"github.com/kballard/go-shellquote"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
)

// recursive function to calculate whole process+childs memory usage
func totalMemoryUsage(p *process.Process) (uint64, error) {
	currentMem, err := p.MemoryInfo()
	if err != nil {
		return 0, err
	}

	child, err := p.Children()
	if err != nil { // doesn't have child
		return currentMem.RSS, nil
	}

	sum := currentMem.RSS
	for _, children := range child {
		childrenUsage, err := totalMemoryUsage(children)
		if err != nil {
			return 0, err
		}
		sum += childrenUsage
	}

	return sum, nil
}

func monitorMem(p *process.Process, memLimit uint64, result chan uint64) {
	// there is also linux only solution with setrlimit:
	// read `man 2 prlimit`  and
	// https://golang.hotexamples.com/examples/syscall/-/Setrlimit/golang-setrlimit-function-examples.html
	// https://www.quora.com/Computer-Programming/What-is-the-simplest-and-most-accurate-way-to-measure-the-memory-used-by-a-program-in-a-programming-contest-environment/answer/Vivek-Prakash-2
	// the current solution should be platform independant
	// (why windows still exists?)

	// using https://pkg.go.dev/github.com/shirou/gopsutil/process#MemoryInfoStat

	for {
		totalUsingMem, err := totalMemoryUsage(p)
		switch err.(type) {
		case nil: // no error
			break
		case *fs.PathError, syscall.Errno:
			// linux: process in proc not found
			// windows, systemcall not found
			return
		default:
			panic(err)
		}

		print("mem usage (KB) : ")
		println(totalUsingMem / 1024)
		if totalUsingMem > memLimit {
			err := p.Kill()
			if err != nil {
				panic(err)
			}
			result <- totalUsingMem
			return
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func Run(command string, outLimit int, memLimit uint64, timeout time.Duration) (string, error) {
	// which language
	// apply rules
	// compile, print compile errors
	// for: run on all test cases
	//   with tests count and out limit and ...

	// sharif-judge use this:
	// https://github.com/mjnaderi/Sharif-Judge/blob/Version-1/tester/runcode.sh

	memUsage := make(chan uint64)

	words, err := shellquote.Split(command)
	if err != nil {
		return "", MalformedCommandError
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // cleanup resources eventually

	// Create the command with our context
	execCmd := exec.CommandContext(ctx, words[0], words[1:]...)

	stdinWriter, err := execCmd.StdinPipe()
	cobra.CheckErr(err)
	_ = stdinWriter
	//_, err = stdinWriter.Write(testInpData)
	//cobra.CheckErr(err)

	stdoutPipe, err := execCmd.StdoutPipe()
	cobra.CheckErr(err)
	outBuf := make([]byte, outLimit+1)

	stderrPipe, err := execCmd.StderrPipe()
	cobra.CheckErr(err)
	errBuf := make([]byte, outLimit+1)

	// finally start the process!
	err = execCmd.Start()
	cobra.CheckErr(err)

	pid := execCmd.Process.Pid
	process, err := process.NewProcess(int32(pid))
	cobra.CheckErr(err)
	go monitorMem(process, memLimit, memUsage)

	bytesRead, err := stdoutPipe.Read(outBuf)
	if bytesRead == outLimit+1 {
		return "", OutputLimitError
	}
	out := outBuf[:bytesRead]

	err = stdoutPipe.Close()
	cobra.CheckErr(err)

	stderrN, err := stderrPipe.Read(errBuf)
	_ = stderrN
	err = stderrPipe.Close()
	cobra.CheckErr(err)

	// finished flag become true
	// and check for any error
	executeErr := execCmd.Wait()

	//println("stderr:")
	//println(string(errBuf[:stderrN]))

	if ctx.Err() == context.DeadlineExceeded {
		return "", TimedOutError
	} else if ctx.Err() != nil {
		panic(err)
	}

	select {
	case <-memUsage:
		return "", MemoryLimitError
	default:

	}

	if executeErr != nil {
		return "", NonZeroExitError
	}

	outStr := string(out)
	if bytesRead == 0 {
		return outStr, NoOutputError
	}

	return outStr, nil
}
