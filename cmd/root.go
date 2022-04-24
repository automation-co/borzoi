package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "borzoi",
	Short: "Makes it easy to manage your codebase.",
	Long: `Borzoi is a tool that makes it easy to manage your codebase.

It helps you a simple interface to standardize your git repos in the same manner.

For more information, please visit
https://github.com/automation-co/borzoi
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
