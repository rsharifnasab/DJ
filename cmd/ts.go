/*
Copyright © 2022 Roozbeh Sharifnasab rsharifnasab@gmail.com
*/
package cmd

import (
	"github.com/rsharifnasab/DJ/pkg/ts"
	"github.com/spf13/cobra"
)

var tsArgs struct {
	question string
	solution string
	language string
}

// tsCmd represents the student command
var tsCmd = &cobra.Command{
	Use:   "ts",
	Short: "Check syntax for allowed structures and imports",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
students can run and test thier codes via this command`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ts.CheckSource(tsArgs.question, tsArgs.solution, tsArgs.language)
		if err != nil {
			println("error : ")
			println(err.Error())
		} else {
			println("no error")
		}
	},
}

func init() {
	rootCmd.AddCommand(tsCmd)
	var err error

	tsCmd.PersistentFlags().StringVarP(
		&tsArgs.solution, "solution", "s",
		"", "Root of your solution directory")
	err = tsCmd.MarkPersistentFlagRequired("solution")
	cobra.CheckErr(err)
	err = tsCmd.MarkPersistentFlagDirname("solution")
	cobra.CheckErr(err)

	tsCmd.PersistentFlags().StringVarP(
		&tsArgs.question, "question", "q",
		"", "the question path that you are answering to")
	err = tsCmd.MarkPersistentFlagRequired("question")
	cobra.CheckErr(err)

	tsCmd.PersistentFlags().StringVarP(
		&tsArgs.language, "language", "l",
		"", "your code's language")
}
