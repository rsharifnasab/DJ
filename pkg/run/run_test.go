package run

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 	run.Run(`echo "hello world"`, outLimit int, memLimit uint64, timeout time.Duration) (string, error) {

func TestSimple(t *testing.T) {
	out, err := Run(`bash -c 'echo "hello world"'`, 5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.Nil(t, err)
	assert.Equal(t, "hello world\n", out)
}

func TestMalformedQoute(t *testing.T) {
	_, err := Run(`bash -c 'echo "hello world"`, 5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.EqualValues(t, MalformedCommandError, err)
}

func TestTimeLimit(t *testing.T) {
	_, err := Run(`bash -c 'while true; do true; done'`, 5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.EqualValues(t, TimedOutError, err)
}

func TestMemoryLimit(t *testing.T) {
	//_, err := Run(`bash -c 'echo "ok"; tail /dev/zero > /dev/null'`, 5*1024, 50*1024, 20*time.Second)
	_, err := Run(`bash -c 'echo "ok"; cat /dev/zero | head -c 150m > /dev/null'`, 5*1024, 50*1024, 10*time.Second)
	// cat /dev/zero | head -c 5G | tail
	assert.EqualValues(t, MemoryLimitError, err)
}

func TestOutputLimit(t *testing.T) {
	_, err := Run(`bash -c 'while true; do echo "text text text"; done'`, 500, 50*1024*1024, 10*time.Second)
	assert.EqualValues(t, OutputLimitError, err)
}

func TestNoOutput(t *testing.T) {
	_, err := Run(`bash -c 'true'`, 5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.EqualValues(t, NoOutputError, err)
}

func TestNonZero(t *testing.T) {
	_, err := Run(`bash -c 'false'`, 5*1024, 50*1024*1024, 100*time.Millisecond)
	assert.EqualValues(t, NonZeroExitError, err)
}
