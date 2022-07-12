package judge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var examplePath string = "../../examples"
var judger1Path string = examplePath + "/judger1"

func TestCheckReq(t *testing.T) {
	checkReq(judger1Path)
}

func TestNoOfTest(t *testing.T) {
	n := numberOfTests(judger1Path)
	assert.Equal(t, 5, n)
}
