package lib

import (
	"fmt"
	"os"

	"github.com/automation-co/borzoi/internal/config"
	"github.com/go-git/go-git/v5"
)

func Clone() {

	fmt.Println("Cloning the repositories...")

	// Read the config file
	conf := config.ReadConfig()

	// Get the path of the current directory
	path := os.Getenv("PWD")

	// Iterate over the repositories
	iterate(conf, path)

}

// -----------------------------------------------------------------------------

func iterate(conf map[string]interface{}, path string) {
	// Iterate over the repositories
	for title, body := range conf {

		// Checks
		if title == "repo" {
			break
		}

		// Check if the directory exists and if not create it
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// Create the directory
			os.Mkdir(path, 0777)
		}

		if body == nil {
			break
		} else {

			// if the body is a map, that has a key "repo" then we have a
			// repository and we can clone it
			if body.(map[string]interface{})["repo"] != nil {

				// Clone the repository at the path
				fmt.Println(
					"Cloning " +
						body.(map[string]interface{})["repo"].(string) +
						" repo in dir " + path,
				)

				_, err := git.PlainClone(path+"/"+title, false, &git.CloneOptions{
					URL:               body.(map[string]interface{})["repo"].(string),
					RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Repository cloned")
				}

			}

			newPath := path + "/" + title

			iterate(body.(map[string]interface{}), newPath)

		}

	}
}

// -----------------------------------------------------------------------------
