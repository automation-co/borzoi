package cmd

import (
	"github.com/automation-co/borzoi/internal/lib"
	"github.com/spf13/cobra"
)

var (
	AccessToken string
	Username    string
	cloneCmd    = &cobra.Command{
		Use:   "clone",
		Short: "Clones the repos",
		Long:  `Clones the repositories in the given directory`,
		Run: func(cmd *cobra.Command, args []string) {
			lib.Clone(Username, AccessToken)
		},
	}
)

func init() {
	rootCmd.AddCommand(cloneCmd)
	cloneCmd.Flags().StringVarP(&AccessToken, "access", "a", "", "personal access token")
	cloneCmd.Flags().StringVarP(&Username, "user", "u", "", "username")
}
