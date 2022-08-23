package student

import (
	"fmt"

	"github.com/rsharifnasab/DJ/pkg/judge"
)

func Run(submission *judge.Submission) {
	submResult := submission.RunSuite()
	fmt.Printf("%s\n", submResult.String())
	fmt.Printf("result dir: %v\n", submission.ResultDir)
}
