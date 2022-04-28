package lib

import (
	"fmt"
	"sync"

	"github.com/automation-co/borzoi/internal/config"
	"github.com/go-git/go-git/v5"
)

// =============================================================================

// Clones the repos in the given config file
func Clone() {

	fmt.Println("Cloning the repositories...")

	// Read the config file
	conf := config.ReadConfig()

	var wg sync.WaitGroup = sync.WaitGroup{}

	// Iterate over the repos in the config file
	for path, url := range conf {
		wg.Add(1)
		go func(url interface{}, path string) {
			// Get the url of the repo
			repoUrl := url.(string)

			// Clone the repo

			_, err := git.PlainClone(path, false, &git.CloneOptions{
				URL:               repoUrl,
				RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			})
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(url, path)

	}
	wg.Wait()
}

// =============================================================================
