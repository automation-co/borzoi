package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/automation-co/borzoi/internal/config"
	"github.com/automation-co/borzoi/internal/types"
	"github.com/automation-co/borzoi/internal/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/manifoldco/promptui"
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

func FreezeClone(username string, accessToken string, privateKeyFile string) {

	fmt.Println("Cloning the state from borzoi-lock file...")

	conf := config.ReadLockFile()

	// Inputting password if privatekeyfile provided
	var password string
	if privateKeyFile != "" {

		prompt := promptui.Prompt{
			Label: "Password",
			Mask:  '*',
		}

		result, err := prompt.Run()
		password = result
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
	}

	// Get username
	usernameLocal := utils.GetUsername()
	if username == "" {
		username = usernameLocal
	}

	// Create waitgroup
	var wg sync.WaitGroup = sync.WaitGroup{}

	for path, meta := range conf {
		wg.Add(1)
		go func(path string, meta interface{}) {

			// Get the metadata of the repo
			m := types.Meta{}

			err := mapstructure.Decode(meta, &m)
			if err != nil {
				panic(err)
			}

			fmt.Printf("  [x]  Cloning %s\n", m.Repo)

			referenceName := "refs/heads/" + m.Branch

			// Clone the repo
			_, err = git.PlainClone(path, false, &git.CloneOptions{
				URL:               m.Repo,
				RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
				SingleBranch:      true,
				ReferenceName:     plumbing.ReferenceName(referenceName),
			})
			if err != nil {

				if err.Error() == "authentication required" {
					_, err := git.PlainClone(path, false, &git.CloneOptions{
						URL: m.Repo,
						Auth: &http.BasicAuth{
							Username: username,
							Password: accessToken, // personal access token
							// needs to be created using github api
						},
						RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
						SingleBranch:      true,
						ReferenceName:     plumbing.ReferenceName(referenceName),
					})
					if err != nil {
						if err.Error() == "repository already exists" {
							fmt.Println("  [o]  Skipping " + path + " because it already exists")
						} else {
							panic(err)
						}
					}

				} else if err.Error() == "error creating SSH agent: \"SSH agent requested, but could not detect Pageant or Windows native SSH agent\"" {
					_, err := os.Stat(privateKeyFile)
					if err != nil {
						fmt.Printf("read file %s failed %s\n", privateKeyFile, err.Error())
						return
					}
					// TODO: make public keys work
					publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyFile, password)
					if err != nil {
						fmt.Printf("generate publickeys failed: %s\n", err.Error())
						return
					}
					auth := publicKeys
					_, err = git.PlainClone(path, false, &git.CloneOptions{
						URL:               m.Repo,
						Auth:              auth,
						RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
						SingleBranch:      true,
						ReferenceName:     plumbing.ReferenceName(referenceName),
					})
					if err != nil {
						if err.Error() == "repository already exists" {
							fmt.Println("  [o]  Skipping " + path + " because it already exists")
						} else {
							fmt.Println("dint do shit")
							panic(err)
						}
					}
				} else if err.Error() == "repository already exists" {
					fmt.Println("  [o]  Skipping " + path + " because it already exists")
				} else {
					panic(err)
				}
			}

			// Reset hard
			err = utils.ResetHard(path, m.Commit)
			if err != nil {
				panic(err)
			}

			wg.Done()

		}(path, meta)

	}

	wg.Wait()

	fmt.Println("")
	fmt.Println("Woof 👍")

}

// =============================================================================
