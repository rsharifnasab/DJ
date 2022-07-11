package run

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 	run.Run(`echo "hello world"`, outLimit int, memLimit uint64, timeout time.Duration) (string, error) {

func Test1(t *testing.T) {
	out, err := Run(`sh -c 'echo "hello world"'`, 5*1024, 50*1024*1024, 1*time.Second)
	assert.Nil(t, err)
	assert.Equal(t, "hello world\n", out)

}

func Test2(t *testing.T) {

}
