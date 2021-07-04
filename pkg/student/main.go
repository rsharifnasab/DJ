package student

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

func Run() {
	const submissionFile = "./examples/solution.cpp"
	_ = submissionFile
	// todo compile

	const compiledFile = "./examples/a.out"

	const timeout = 2 * time.Second

	const testsCount = 7

	const outLimit int64 = /* 1024 * 1024 * */ 10

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

		//if err != nil {
		//	panic(err)
		//}

		//println("stdout pipe created")
		limitedReader := &io.LimitedReader{R: stdoutPipe, N: outLimit}
		//println("limited reader created")
		out, err := ioutil.ReadAll(limitedReader)
		//println("after read all")

		err = cmd.Wait()
		//println("wait complete")
		/*
			 //out, err := cmd.Output()
			OR
			   limitedReader := &io.LimitedReader{R: response.Body, N: limit}
			   body, err := ioutil.ReadAll(limitedReader)

			   or

			   body, err := ioutil.ReadAll(io.LimitReader(response.Body, limit))
		*/

		//cancel()
		//time.Sleep(3 * time.Second)
		// here should not have performace critical code

		//err = cmd.Run()
		//if ctx.Err() != nil {
		//	print(".")
		//	println(ctx.Err().Error()) // it says context cancelled
		//}
		//cmd.Wait()

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
