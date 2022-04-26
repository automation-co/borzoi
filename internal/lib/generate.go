package lib

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/automation-co/borzoi/internal/utils"
)

func Generate() {
	fmt.Println("Generating the config file...")

	// type Repo struct {
	// 	Name string
	// 	URL  string
	// }

	// type Folder struct {
	// 	Name    string
	// 	Repos   []Repo
	// 	Folders []Folder
	// }

	// folders := Folder{}

	// folders := make(map[string]interface{})

	// Recursing over the directories in the current directory
	err := filepath.WalkDir(
		".",
		func(path string, d os.DirEntry, err error) error {

			if err != nil {
				return err
			}

			// Stuff goes here...
			// fmt.Println(path)

			// check if the directory is a git repository
			// if it is, then we can add it to the config file as a repository
			// if it is not, then we can add it to the config file as a folder

			isGitRepo := utils.IsGitRepo(path)

			if isGitRepo {

				fmt.Println("Git repo found : " + path)
			}

			return nil

		},
	)
	if err != nil {
		panic(err)
	}

}
