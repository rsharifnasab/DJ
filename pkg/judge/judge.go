package judge

import (
	"regexp"
	"strconv"
	"strings"

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

func judge(judgerPath string) {
	checkReq(judgerPath)

	noOfTests := numberOfTests(judgerPath)
	println(noOfTests)
}
