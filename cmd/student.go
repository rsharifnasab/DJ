/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/rsharifnasab/DJ/pkg/run"
	"github.com/rsharifnasab/DJ/pkg/student"
	"github.com/spf13/cobra"
)

var (
	submission run.Submission
)

// studentCmd represents the student command
var studentCmd = &cobra.Command{
	Use:   "student",
	Short: "students will use this sub-command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

students can run and test thier codes via this command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("student cmd")
		student.Run(submission)
	},
}

func init() {
	rootCmd.AddCommand(studentCmd)
	var err error

	// Persistent Flags which will work for this command and all subcommands
	studentCmd.PersistentFlags().StringVarP(
		&submission.Path, "submission", "s",
		"", "Root of submission directory")
	err = studentCmd.MarkPersistentFlagRequired("submission")
	cobra.CheckErr(err)
	err = studentCmd.MarkPersistentFlagDirname("submission")
	cobra.CheckErr(err)

	studentCmd.PersistentFlags().StringVarP(
		&submission.Judger, "judger", "j",
		"", "The judger suitable for your submission")
	err = studentCmd.MarkPersistentFlagRequired("judger")
	cobra.CheckErr(err)

	studentCmd.PersistentFlags().StringVarP(
		&submission.Question, "question", "q",
		"", "the question you are answering to")

}
