/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rsharifnasab/DJ/pkg/judge"
	"github.com/rsharifnasab/DJ/pkg/student"
	"github.com/spf13/cobra"
)

var (
	submission judge.Submission
)

// studentCmd represents the student command
var studentCmd = &cobra.Command{
	Use:   "student",
	Short: "students will use this sub-command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

students can run and test thier codes via this command`,
	Run: func(cmd *cobra.Command, args []string) {
		student.Run(&submission)
	},
}

func init() {
	rootCmd.AddCommand(studentCmd)
	var err error

	// Persistent Flags which will work for this command and all subcommands
	studentCmd.PersistentFlags().StringVarP(
		&submission.UserSolution, "solution", "s",
		"", "Root of your solution directory")
	err = studentCmd.MarkPersistentFlagRequired("solution")
	cobra.CheckErr(err)
	err = studentCmd.MarkPersistentFlagDirname("solution")
	cobra.CheckErr(err)

	studentCmd.PersistentFlags().StringVarP(
		&submission.Judger, "judger", "j",
		"", "The judger suitable for your submission")
	err = studentCmd.MarkPersistentFlagRequired("judger")
	cobra.CheckErr(err)

	studentCmd.PersistentFlags().StringVarP(
		&submission.Question, "question", "q",
		"", "the question path that you are answering to")
	err = studentCmd.MarkPersistentFlagRequired("question")
	cobra.CheckErr(err)

	studentCmd.PersistentFlags().StringVarP(
		&submission.Result, "result", "r",
		"", "where to save result and logs")

	studentCmd.PersistentFlags().StringVarP(
		&submission.Language, "language", "l",
		"", "your code's language")
}
