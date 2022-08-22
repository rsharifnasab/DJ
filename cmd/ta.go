package cmd

import (
	"fmt"

	"github.com/rsharifnasab/DJ/pkg/ta"
	"github.com/spf13/cobra"
)

var (
	taSubmission ta.TaSubmission
)

// taCmd represents the ta command
var taCmd = &cobra.Command{
	Use:   "ta",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ta called")
	},
	Hidden: true,
}

func init() {
	rootCmd.AddCommand(taCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	var err error

	// Persistent Flags which will work for this command and all subcommands
	taCmd.PersistentFlags().StringVarP(
		&submission.Solution, "solution", "s",
		"", "Root of your solution directory")
	err = taCmd.MarkPersistentFlagRequired("solution")
	cobra.CheckErr(err)
	err = taCmd.MarkPersistentFlagDirname("solution")
	cobra.CheckErr(err)

	taCmd.PersistentFlags().StringVarP(
		&submission.Judger, "judger", "j",
		"", "The judger suitable for your submission")
	err = taCmd.MarkPersistentFlagRequired("judger")
	cobra.CheckErr(err)

	taCmd.PersistentFlags().StringVarP(
		&submission.Question, "question", "q",
		"", "the question path that you are answering to")
	err = taCmd.MarkPersistentFlagRequired("question")
	cobra.CheckErr(err)

}
