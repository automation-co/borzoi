package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the config file",
	Long:  `Generates the config file for borzoi using your current file structure.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
