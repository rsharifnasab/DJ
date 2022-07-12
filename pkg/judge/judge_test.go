package judge

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
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

func TestJudge(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "dj_judge_*")
	cobra.CheckErr(err)
	defer os.RemoveAll(tempDir)

	//judge(judger1Path, )
}
