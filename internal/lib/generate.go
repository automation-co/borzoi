package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/automation-co/borzoi/internal/utils"
)

// =============================================================================

func Generate() {

	repos := make(map[string]interface{})

	// Recursing over the directories in the current directory
	err := filepath.WalkDir(
		".",
		func(path string, d os.DirEntry, err error) error {

			if err != nil {
				return err
			}

			isIgnored := utils.IsIgnored(path)

			if isIgnored {
				return filepath.SkipDir
			}

			isGitRepo := utils.IsGitRepo(path)

			if isGitRepo {
				// Get the url of the repo
				url, err := utils.GetRepoUrl(path)

				if err != nil {
					return err
				}
				// add the repo to the repos map
				repos[path] = url
				return filepath.SkipDir
			}

			return nil

		},
	)
	if err != nil {
		panic(err)
	}

	// Write the config file ---------------------------------------------------

	// convert the repos to a json string
	jsonString, err := json.Marshal(repos)

	// write the file as borzoi.json
	err = ioutil.WriteFile("borzoi.json", jsonString, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Config file generated 👍")

	// -------------------------------------------------------------------------

}

// =============================================================================
