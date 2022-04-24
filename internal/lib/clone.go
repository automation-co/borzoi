package lib

import (
	"fmt"
	"os"

	"github.com/automation-co/borzoi/internal/config"
)

// TASKS TODO:
// // 1. Handle the case where the config file is empty
// // 2. Handle the path management
// 3. Clone the repositories and its error handling

func iterate(conf map[string]interface{}, path string) {
	// Iterate over the repositories
	for title, body := range conf {

		// Checks
		if title == "repo" {
			break
		}

		newPath := path + "/" + title
		fmt.Println(newPath)

		if body == nil {
			break
		} else {

			// if the body is a map, that has a key "repo" then we have a repository
			// and we can clone it so for now lets print the value of the key "repo"
			if body.(map[string]interface{})["repo"] != nil {
				fmt.Println(body.(map[string]interface{})["repo"])
			}

			iterate(body.(map[string]interface{}), newPath)

		}

	}
}

func Clone() {

	fmt.Println("Cloning the repositories...")

	// Read the config file
	conf := config.ReadConfig()

	// Print the config file contents
	fmt.Println("Config file contents:")
	// fmt.Println(conf)

	// Get the path of the current directory
	path := os.Getenv("PWD")

	// Iterate over the repositories
	iterate(conf, path)

}
