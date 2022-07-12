package run

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	out, _, err := Run(`bash -c 'echo "hello world"'`, 5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.Nil(t, err)
	assert.Equal(t, "hello world\n", out)
}

func TestStderr(t *testing.T) {
	stdout, stderr, err := Run(`bash -c '1>&2 echo "hello err"; echo "hello out"'`, 5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.Nil(t, err)
	assert.Equal(t, "hello err\n", stderr)
	assert.Equal(t, "hello out\n", stdout)
}

func TestMalformedQoute(t *testing.T) {
	_, _, err := Run(`bash -c 'echo "hello world"`, 5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.EqualValues(t, MalformedCommandError, err)
}

func TestTimeLimit(t *testing.T) {
	_, _, err := Run(`bash -c 'while true; do true; done'`, 5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.EqualValues(t, TimedOutError, err)
}

func TestMemoryLimit(t *testing.T) {
	_, _, err := Run(`bash -c 'echo "ok"; cat /dev/zero | head -c 500m > /dev/null'`, 5*1024, 50*1024, 10*time.Second)
	assert.EqualValues(t, MemoryLimitError, err)
}

func TestOutputLimit(t *testing.T) {
	_, _, err := Run(`bash -c 'while true; do echo "text text text"; done'`, 512, 50*1024*1024, 10*time.Second)
	assert.NotNil(t, err)
	assert.Contains(t, []error{OutputLimitError, NonZeroExitError}, err)
	// why non-zero?
}

func TestNoOutput(t *testing.T) {
	_, _, err := Run(`bash -c 'true'`, 5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.EqualValues(t, NoOutputError, err)
}

func TestNonZero(t *testing.T) {
	_, _, err := Run(`bash -c 'false'`, 5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.EqualValues(t, NonZeroExitError, err)
}

func TestNonExistingPath(t *testing.T) {
	_, _, err := Run(`./non_existing_file`,
		5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.EqualValues(t, NotValidExecutableError, err)
}

func TestNonExecutable(t *testing.T) {
	file, err := ioutil.TempFile("", "script*.sh")
	assert.Nil(t, err)
	defer os.Remove(file.Name())

	content := `#!/usr/bin/env bash 
echo "hello world"
`
	_, err = file.Write([]byte(content))
	assert.Nil(t, err)

	_, _, err = Run(file.Name(),
		5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.EqualValues(t, NotValidExecutableError, err)
}

func TestExecutable(t *testing.T) {
	file, err := ioutil.TempFile("", "script*.sh")
	assert.Nil(t, err)
	defer os.Remove(file.Name())

	content := `#!/usr/bin/env bash 
echo "hello world"
`
	_, err = file.Write([]byte(content))
	assert.Nil(t, err)

	file.Chmod(0777)
	err = file.Close()
	assert.Nil(t, err)

	stdout, _, err := Run(file.Name(),
		5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.Nil(t, err)
	assert.EqualValues(t, "hello world\n", stdout)
}
