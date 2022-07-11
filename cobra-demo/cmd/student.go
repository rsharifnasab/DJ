/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// studentCmd represents the student command
var studentCmd = &cobra.Command{
	Use:   "student",
	Short: "students will use this sub-command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

students can run and test thier codes via this command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "\n----\n%+v\n----\n", cmd.Flags())

		fmt.Println("student called")
	},
}

func init() {
	rootCmd.AddCommand(studentCmd)


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// studentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// studentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
