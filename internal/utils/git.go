package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/automation-co/borzoi/internal/types"
	"github.com/go-git/go-git/v5/config"
)

// -----------------------------------------------------------------------------

// isDirectory determines if a file represented
// by `path` is a directory or not
func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

// -----------------------------------------------------------------------------

// Tells if the directory is a git repo
func IsGitRepo(path string) bool {

	path, err := filepath.Abs(path + "/.git")
	if err != nil {
		panic(err)
	}

	if condition, err := IsDirectory(path); condition && err == nil {
		return true
	}

	return false
}

// -----------------------------------------------------------------------------

// Get repo meta data
func GetRepoMetaData(path string) (types.Meta, error) {

	var meta types.Meta

	pathOfGitConfig := path + "/.git/config"

	gitConfig, err := ioutil.ReadFile(pathOfGitConfig)
	if err != nil {
		panic(err)
	}

	gitConfigString := string(gitConfig)
	reader := strings.NewReader(gitConfigString)
	goGitConfig, err := config.ReadConfig(reader)
	if err != nil {
		return types.Meta{}, err
	}

	urls := goGitConfig.Remotes["origin"].URLs
	if len(urls) < 1 {
		return types.Meta{}, fmt.Errorf("no remote origin found")
	}

	url := urls[0]

	meta.Repo = url

	// get the branch name
	pathOfHead := path + "/.git/HEAD"
	head, err := ioutil.ReadFile(pathOfHead)
	if err != nil {
		return types.Meta{}, err
	}

	headString := string(head)
	headString = strings.TrimSpace(headString)
	headString = strings.TrimPrefix(headString, "ref: refs/heads/")
	headString = strings.TrimSuffix(headString, "")

	meta.Branch = headString

	// get the commit hash

	pathOfCommit := path + "/.git/" + "refs/heads/" + headString

	commit, err := ioutil.ReadFile(pathOfCommit)
	if err != nil {
		return types.Meta{}, err
	}

	commitString := string(commit)
	commitString = strings.TrimSpace(commitString)
	commitString = strings.TrimPrefix(commitString, "")
	commitString = strings.TrimSuffix(commitString, "")

	meta.Commit = commitString

	return meta, nil

}

// -----------------------------------------------------------------------------

// Returns the url of the git repo
func GetRepoUrl(path string) (string, error) {
	pathOfGitConfig := path + "/.git/config"

	gitConfig, err := ioutil.ReadFile(pathOfGitConfig)
	if err != nil {
		panic(err)
	}

	gitConfigString := string(gitConfig)
	reader := strings.NewReader(gitConfigString)
	goGitConfig, err := config.ReadConfig(reader)
	if err != nil {
		return "", err
	}

	urls := goGitConfig.Remotes["origin"].URLs

	if len(urls) < 1 {
		return "", fmt.Errorf("no remote origin found")
	}

	url := urls[0]

	return url, err
}

// Tells if the file needs to be ignored
func IsIgnored(path string) bool {
	readFile, err := os.Open(".borzoiignore")

	if err != nil {
		return false
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		ignoreQuery := fileScanner.Text()
		if strings.Contains(path, ignoreQuery) {
			return true
		}
	}

	readFile.Close()
	return false
}

// Gets global config of Git
func GetUsername() string {

	// get the path of the global git config
	gitConfig, err := config.LoadConfig(config.GlobalScope)
	if err != nil {
		panic(err)
	}
	// convert to json
	gitConfigJson, err := gitConfig.Marshal()
	if err != nil {
		panic(err)
	}

	gitConfigString := string(gitConfigJson)
	reader := strings.NewReader(gitConfigString)
	goGitConfig, err := config.ReadConfig(reader)
	if err != nil {
		panic(err)
	}

	username := goGitConfig.User.Name

	return username

}

// -----------------------------------------------------------------------------
