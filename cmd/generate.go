package cmd

import (
	"github.com/automation-co/borzoi/internal/lib"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the config file",
	Long:  `Generates the config file for borzoi using your current file structure.`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.Generate()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
