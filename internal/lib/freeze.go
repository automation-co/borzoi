package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/automation-co/borzoi/internal/config"
	"github.com/automation-co/borzoi/internal/types"
	"github.com/automation-co/borzoi/internal/utils"
	"github.com/mitchellh/mapstructure"
)

// =============================================================================

func Freeze() {

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
				meta, err := utils.GetRepoMetaData(path)

				if err != nil {
					return err
				}
				// add the repo to the repos map
				repos[path] = meta
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
	err = ioutil.WriteFile("borzoi-lock.json", jsonString, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Lock file generated 👍")

	// -------------------------------------------------------------------------

}

// =============================================================================

func FreezeClone() {

	fmt.Println("Cloning the state from borzoi-lock file...")

	conf := config.ReadLockFile()

	for path, meta := range conf {
		fmt.Println(path)

		m := types.Meta{}

		err := mapstructure.Decode(meta, &m)

		if err != nil {
			panic(err)
		}

		fmt.Println("repo : ", m.Repo)
		fmt.Println("branch : ", m.Branch)
		fmt.Println("commit : ", m.Commit)

	}

}

// =============================================================================
