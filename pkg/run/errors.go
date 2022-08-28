package run

import "fmt"

var (
	ErrNonZeroExit        = fmt.Errorf("nonZero Exit: The code finished with exit code != 0")
	ErrOutputLimit        = fmt.Errorf("output Limit: Produced output is more that expected")
	ErrMemoryLimit        = fmt.Errorf("memory Limit: The code is using too much memory")
	ErrTimedOut           = fmt.Errorf("timed Out: the code did'nt complete execution in time")
	ErrNoOutput           = fmt.Errorf("no Output: The Code doesn't have any output")
	ErrMalformedCommand   = fmt.Errorf("malformed Command: The provided command is malformed")
	ErrNotValidExecutable = fmt.Errorf("specified program is not a valid executable")
)

// with the help of: https://www.domjudge.org/docs/manual/8.0/team.html

//TODO: replace with proper error structs
