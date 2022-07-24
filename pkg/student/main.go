package student

import (
	"github.com/rsharifnasab/DJ/pkg/judge"
	"github.com/rsharifnasab/DJ/pkg/util"
)

func Run(submission *judge.Submission) {
	util.PrintStruct(judge.RunSubmission(submission))
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
		//println(submission.SourceContent)
		util.PrintStruct(question.AvailableLangs[submission.LanguageName])
		println("\n----------\n")

		runExampleTests()

}
*/
