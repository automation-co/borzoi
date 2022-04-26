package utils

import (
	"os"
	"path/filepath"
)

// isDirectory determines if a file represented
// by `path` is a directory or not
func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

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
