/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rsharifnasab/DJ/pkg/run"
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
		fmt.Fprintf(os.Stdout, "---\n%+v\n----\n", submission)

		fmt.Println("student called")

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
		&submission.Runner, "language", "l",
		"", "Language of your submission")
	err = studentCmd.MarkPersistentFlagRequired("language")
	cobra.CheckErr(err)
}
