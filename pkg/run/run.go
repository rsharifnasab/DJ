package run

import (
	"context"
	"fmt"
	"io/fs"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/kballard/go-shellquote"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		// TODO: check, what if error in child?
		childrenUsage, err := totalMemoryUsage(children)
		if err != nil {
			return 0, err
		}
		sum += childrenUsage
	}

	return sum, nil
}

// Frequently monitor memory usage of given process.
// if is it using memory more than expected? kill it and write somethin to result channel
// using https://pkg.go.dev/github.com/shirou/gopsutil/process#MemoryInfoStat
// the current solution should work on windows too
func monitorMem(pid int, memLimit uint64, result chan uint64) {
	// there is also linux only solution with setrlimit:
	// read `man 2 prlimit`  and
	// https://golang.hotexamples.com/examples/syscall/-/Setrlimit/golang-setrlimit-function-examples.html
	// https://www.quora.com/Computer-Programming/What-is-the-simplest-and-most-accurate-way-to-measure-the-memory-used-by-a-program-in-a-programming-contest-environment/answer/Vivek-Prakash-2
	// (why windows still exists?)

	process, err := process.NewProcess(int32(pid))
	cobra.CheckErr(err)

	for {
		totalUsingMem, err := totalMemoryUsage(process)
		switch err.(type) {
		case nil: // no error, break the switch case
			break
		case *fs.PathError, syscall.Errno:
			// linux: process in proc not found
			// windows, systemcall error
			return
		default:
			panic(err)
		}

		if totalUsingMem > memLimit {
			err := process.Kill()
			if err != nil {
				panic(err)
			}
			result <- totalUsingMem
			return
		}

		time.Sleep(50 * time.Millisecond)
	}
}

// run a command with given limits for outbut (bytes), memory limiy (bytes) and duration
// duration is handling with golang's context so it's almost reliable
// but output limit and memory limit are handmaiden cross platform solutions
// known bugs: memory limit monitor routine sometimes experience starvation so it doen'st kill the program on-time
func Run(commandStr string, outLimit int, memLimit uint64, timeout time.Duration) (string, string, error) {
	// sharif-judge use this:
	// https://github.com/mjnaderi/Sharif-Judge/blob/Version-1/tester/runcode.sh

	memUsageResult := make(chan uint64)

	if viper.GetBool("debug") {
		fmt.Println(commandStr)
	}

	commandWords, err := shellquote.Split(commandStr)
	if err != nil {
		return "", "", MalformedCommandError
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // cleanup resources eventually

	// Create the command with our context
	execCmd := exec.CommandContext(ctx, commandWords[0], commandWords[1:]...)

	// initialize stdin, but we don't write anything by now
	stdinWriter, err := execCmd.StdinPipe()
	if err != nil {
		print("err : ")
		fmt.Println(err.Error())
		//cobra.CheckErr(err)
		panic(err)
	}
	_ = stdinWriter
	//_, err = stdinWriter.Write(testInpData)
	//cobra.CheckErr(err)

	// initialize stdout and stderr before start
	stdoutPipe, err := execCmd.StdoutPipe()
	cobra.CheckErr(err)
	outBuf := make([]byte, outLimit+1)

	stderrPipe, err := execCmd.StderrPipe()
	cobra.CheckErr(err)
	errBuf := make([]byte, outLimit+1)

	// finally start the process!
	err = execCmd.Start()
	if err != nil {
		switch err.(type) {
		case *fs.PathError:
			//fmt.Println(err.Error())
			return "", "", NotValidExecutableError

		default:
			cobra.CheckErr(err)
		}
	}

	pid := execCmd.Process.Pid
	go monitorMem(pid, memLimit, memUsageResult)

	// fill stdout buffer
	bytesRead, err := stdoutPipe.Read(outBuf)
	if bytesRead == outLimit+1 {
		return "", "", OutputLimitError
	}

	err = stdoutPipe.Close()
	cobra.CheckErr(err)

	// fill stderr buffer
	stderrN, err := stderrPipe.Read(errBuf)
	_ = stderrN
	err = stderrPipe.Close()
	cobra.CheckErr(err)

	// finished flag become true
	// and check for any error
	executeErr := execCmd.Wait()

	if ctx.Err() == context.DeadlineExceeded {
		return "", "", TimedOutError
	} else if ctx.Err() != nil {
		panic(err)
	}

	select {
	case <-memUsageResult:
		return "", "", MemoryLimitError
	default:

	}

	outStr := string(outBuf[:bytesRead])
	errStr := string(errBuf[:stderrN])

	if executeErr != nil {
		fmt.Println(executeErr.Error()) // TODO: handle in a better way
		return outStr, errStr, NonZeroExitError
	}

	if bytesRead == 0 {
		return outStr, errStr, NoOutputError
	}

	return outStr, errStr, nil
}

func DefaultRun(command string) (string, string, error) {
	stdout, stderr, err := Run(command, 10*1024*1024, 1024*1024*1024, 20*time.Second)
	return stdout, stderr, err
}

func JustOut(command string) (string, error) {
	stdout, stderr, err := DefaultRun(command)
	if len(stderr) > 0 {
		stderr = strings.TrimSpace(stderr)
		fmt.Printf("STDERR: [%v]\n", stderr)
	}
	return stdout, err
}

func JustRun(command string) error {
	stdout, stderr, err := DefaultRun(command)
	if len(stderr) > 0 {
		stderr = strings.TrimSpace(stderr)
		fmt.Printf("STDERR: [%v]\n", stderr)
	}
	if len(stdout) > 0 {
		stdout = strings.TrimSpace(stdout)
		fmt.Printf("STDOUT: [%v]\n", stdout)
	}
	return err
}
