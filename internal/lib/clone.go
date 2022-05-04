package lib

import (
	"fmt"
	"sync"

	"github.com/automation-co/borzoi/internal/config"
	"github.com/automation-co/borzoi/internal/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// =============================================================================

// Clones the repos in the given config file
func Clone(username string, accessToken string) {

	fmt.Println("Cloning the repositories...")
	fmt.Println("")

	// Read the config file
	conf := config.ReadConfig()

	// Get username
	usernameLocal := utils.GetUsername()
	if username == "" {
		username = usernameLocal
	}

	// Create waitgroup
	var wg sync.WaitGroup = sync.WaitGroup{}

	// Iterate over the repos in the config file
	for path, url := range conf {
		wg.Add(1)
		go func(url interface{}, path string) {
			// Get the url of the repo
			repoUrl := url.(string)

			fmt.Printf("  [x]  Cloning %s\n", repoUrl)

			// Clone the repo
			_, err := git.PlainClone(path, false, &git.CloneOptions{
				URL:               repoUrl,
				RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			})
			if err != nil {

				if err.Error() == "authentication required" {
					_, err := git.PlainClone(path, false, &git.CloneOptions{
						URL: repoUrl,
						Auth: &http.BasicAuth{
							Username: username,
							Password: accessToken, // personal access token
							// needs to be created using github api
						},
						RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
					})
					if err != nil {
						if err.Error() == "repository already exists" {
							fmt.Println("  [o]  Skipping " + path + " because it already exists")
						} else {
							panic(err)
						}
					}
				} else if err.Error() == "repository already exists" {
					fmt.Println("  [o]  Skipping " + path + " because it already exists")
				} else {
					panic(err)
				}
			}
			wg.Done()
		}(url, path)

	}
	wg.Wait()

	fmt.Println("")
	fmt.Println("Woof 👍")
}

// =============================================================================
