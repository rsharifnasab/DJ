package run

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	out, _, err := DefaultRun(`bash -c 'echo "hello world"'`)
	assert.NoError(t, err)
	assert.Equal(t, "hello world\n", out)
}

// TODO: refactor run with justTun, so there are no extra args

func TestStderr(t *testing.T) {
	stdout, stderr, err := DefaultRun(`bash -c '1>&2 echo "hello err"; echo "hello out"'`)
	assert.NoError(t, err)
	assert.Equal(t, "hello err\n", stderr)
	assert.Equal(t, "hello out\n", stdout)
}

func TestMalformedQoute(t *testing.T) {
	_, _, err := DefaultRun(`bash -c 'echo "hello world"`)
	assert.EqualValues(t, ErrMalformedCommand, err)
}

func TestTimeLimit(t *testing.T) {
	_, _, err := Run(`bash -c 'while true; do true; done'`, 5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.EqualValues(t, ErrTimedOut, err)
}

func TestMemoryLimit(t *testing.T) {
	_, _, err := Run(`bash -c 'echo "ok"; cat /dev/zero | head -c 500m > /dev/null'`, 5*1024, 50*1024, 10*time.Second)
	assert.EqualValues(t, ErrMemoryLimit, err)
}

func TestOutputLimit(t *testing.T) {
	_, _, err := Run(`bash -c 'while true; do echo "text text text"; done'`, 100, 50*1024*1024, 5*time.Second)
	assert.Error(t, err)
	assert.Contains(t, []error{ErrOutputLimit, ErrNonZeroExit, ErrTimedOut}, err)
	// why non-zero? wht time limit on windows?
}

func TestNoOutput(t *testing.T) {
	_, _, err := DefaultRun(`bash -c 'true'`)
	assert.EqualValues(t, ErrNoOutput, err)
}

func TestNonZero(t *testing.T) {
	_, _, err := DefaultRun(`bash -c 'false'`)
	assert.EqualValues(t, ErrNonZeroExit, err)
}

func TestNonExistingPath(t *testing.T) {
	_, _, err := DefaultRun(`./non_existing_file`)
	assert.EqualValues(t, ErrNotValidExecutable, err)
}

func TestNonExecutable(t *testing.T) {
	file, err := os.CreateTemp("", "script*.sh")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	content := `#!/usr/bin/env bash
echo "hello world"
`
	_, err = file.Write([]byte(content))
	assert.NoError(t, err)

	_, _, err = DefaultRun(file.Name())
	assert.EqualValues(t, ErrNotValidExecutable, err)
}

func TestExecutable(t *testing.T) {
	file, err := os.CreateTemp("", "script*.sh")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	content := `#!/usr/bin/env bash
echo "hello world"
`
	_, err = file.Write([]byte(content))
	assert.NoError(t, err)

	file.Chmod(0777)
	err = file.Close()
	assert.NoError(t, err)

	stdout, _, err := Run(file.Name(),
		5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.NoError(t, err)
	assert.EqualValues(t, "hello world\n", stdout)
}

func TestChild(t *testing.T) {
	file, err := os.CreateTemp("", "script*.sh")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	content := `#!/usr/bin/env bash	
(sleep 0.2; echo "done") &
sleep 0.5
`
	_, err = file.Write([]byte(content))
	assert.NoError(t, err)

	file.Chmod(0777)
	err = file.Close()
	assert.NoError(t, err)

	stdout, _, err := Run(file.Name(),
		5*1024, 50*1024*1024, 1000*time.Millisecond)
	assert.NoError(t, err)
	assert.EqualValues(t, "done\n", stdout)
}
