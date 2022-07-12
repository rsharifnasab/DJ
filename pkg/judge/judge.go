package judge

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/plus3it/gorecurcopy"
	"github.com/rsharifnasab/DJ/pkg/run"
	"github.com/spf13/cobra"
)

func checkReq(judgerPath string) {

	stdout, err := run.JustOut(judgerPath + "/req.sh")
	cobra.CheckErr(err)
	space := regexp.MustCompile(`\s+`)
	reqWords := space.Split(stdout, -1)
	run.CheckAndErrorRequirements(reqWords)
}

func numberOfTests(judgerPath string) int {
	stdout, err := run.JustOut(judgerPath + "/run.sh count")
	cobra.CheckErr(err)
	n, err := strconv.Atoi(strings.TrimSpace(stdout))
	cobra.CheckErr(err)
	return n
}
func makeTempfolder() string {
	tmpFolder, err := ioutil.TempDir("", "djtmp*")
	cobra.CheckErr(err)
	return tmpFolder
}

func copyLibFiles(judgerPath, tmpFolder string) {
	gorecurcopy.CopyDirectory(judgerPath+"/lib", tmpFolder)
}

func judge(judgerPath string, submissionPath string) {
	checkReq(judgerPath)
	tmpFolder := makeTempfolder()
	copyLibFiles(judgerPath, tmpFolder)

	noOfTests := numberOfTests(judgerPath)
	_ = noOfTests
}
