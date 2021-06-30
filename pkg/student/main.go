package student

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Run() {
	const submissionFile = "/tmp/dj/a.cpp"
	_ = submissionFile
	// todo compile

	const compiledFile = "/tmp/dj/a.out"

	const timeout = 2 * time.Second

	for i := 1; i <= 4; i++ {
		testAddr := fmt.Sprintf("/tmp/dj/tests/in/input%d.txt", i)
		rawFile, err := os.Open(testAddr)
		if err != nil {
			panic(err)
		}

		bufReader := bufio.NewReader(rawFile)

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel() // cleanup resources

		// Create the command with our context
		cmd := exec.CommandContext(ctx, compiledFile)

		stdin, err := cmd.StdinPipe()
		if err != nil {
			panic(err)
		}
		stdinBuf := bufio.NewWriter(stdin)
		stdinBuf.ReadFrom(bufReader)
		if _, err = io.WriteString(stdin, "0\n"); err != nil {
			panic(err)
		}

		err = cmd.Run()

		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("time limit")
			continue
		}

		out, _ := cmd.Output()
		outStr := strings.TrimSpace(string(out))
		fmt.Printf("Output: [%v]\n", outStr)

		if err != nil {
			fmt.Println("Non-zero exit code or program closed", err)
		}
	}
}
