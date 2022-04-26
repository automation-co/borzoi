package lib

import (
	"fmt"

	"github.com/automation-co/borzoi/internal/config"
	"github.com/go-git/go-git/v5"
)

// =============================================================================

// Clones the repos in the given config file
func Clone() {

	fmt.Println("Cloning the repositories...")

	// Read the config file
	conf := config.ReadConfig()

	// Iterate over the repos in the config file
	for path, url := range conf {

		// Get the url of the repo
		url := url.(string)

		// Clone the repo

		_, err := git.PlainClone(path, false, &git.CloneOptions{
			URL:               url,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		})
		if err != nil {
			panic(err)
		}

	}

}

// =============================================================================
