package student

import (
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

type Question struct {
}

type TestCase struct {
	Question
}

func oneTestCase(i int) {

}

func Run() {
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
