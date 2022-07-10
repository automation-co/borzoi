package cmd

import (
	"github.com/automation-co/borzoi/internal/lib"
	"github.com/automation-co/borzoi/internal/utils"
	"github.com/spf13/cobra"
)

var (
	AccessToken    string
	Username       string
	Password       string
	PrivateKeyFile string

	cloneCmd = &cobra.Command{
		Use:   "clone",
		Short: "Clones the repos",
		Long:  `Clones the repositories in the given directory`,
		Run: func(cmd *cobra.Command, args []string) {

			// Check for the lock file
			if utils.IsLockFilePresent() {

				// Check for the config file
				if utils.IsConfigFilePresent() {
					lib.Clone(Username, AccessToken, PrivateKeyFile)
				} else {
					lib.FreezeClone(Username, AccessToken, PrivateKeyFile)
				}

			} else {
				lib.Clone(Username, AccessToken, PrivateKeyFile)
			}

		},
	}
)

func init() {
	rootCmd.AddCommand(cloneCmd)
	cloneCmd.Flags().StringVarP(&AccessToken, "access", "a", "", "personal access token")
	cloneCmd.Flags().StringVarP(&Username, "user", "u", "", "username")
	cloneCmd.Flags().StringVarP(&Password, "password", "p", "", "password")
	cloneCmd.Flags().StringVarP(&PrivateKeyFile, "privateKeyFile", "f", "", "location")
}
