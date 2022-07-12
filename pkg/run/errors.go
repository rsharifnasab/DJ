package run

import "fmt"

var (
	//CompileError     error = fmt.Errorf("Compile Error: The code cound't be compiled.")
	//WrongAnswerError error = fmt.Errorf("No Output: The Code doesn't have any output")
	NonZeroExitError        error = fmt.Errorf("NonZero Exit: The code finished with exit code != 0")
	OutputLimitError        error = fmt.Errorf("Output Limit: Produced output is more that expected.")
	MemoryLimitError        error = fmt.Errorf("Memory Limit: The code is using too much memory.")
	TimedOutError           error = fmt.Errorf("Timed Out: the code did'nt complete execution in time")
	NoOutputError           error = fmt.Errorf("No Output: The Code doesn't have any output")
	MalformedCommandError   error = fmt.Errorf("Malformed Command: The provided command is malformed")
	NotValidExecutableError error = fmt.Errorf("Specified program is not a valid executable")
)

// with the help of: https://www.domjudge.org/docs/manual/8.0/team.html
