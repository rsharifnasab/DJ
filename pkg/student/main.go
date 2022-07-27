package student

import (
	"fmt"

	"github.com/rsharifnasab/DJ/pkg/judge"
)

func Run(submission *judge.Submission) {
	submResult := judge.RunSubmission(submission)
	fmt.Printf("%s\n", submResult.String())
	//TODO: print student details and final score
}

/*
func Run() {

	judge, err := judge.InitJudge("./examples")
	if err != nil {
		panic(err)
	}
	util.PrintStruct(judge)

	question, err := question.NewQuestion("./examples/Q1", judge)
	if err != nil {
		panic(err)
	}
	util.PrintStruct(question)

	/*
		submission := run.NewSubmission("./examples/solution.cpp")

		util.PrintStruct(submission)
		//fmt.Println(submission.SourceContent)
		util.PrintStruct(question.AvailableLangs[submission.LanguageName])
		fmt.Println("\n----------\n")

		runExampleTests()

}
*/
