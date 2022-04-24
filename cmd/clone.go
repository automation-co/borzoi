package cmd

import (
	"github.com/automation-co/borzoi/internal/lib"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clones the repos",
	Long:  `Clones the repositories in the given directory`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.Clone()
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

}
