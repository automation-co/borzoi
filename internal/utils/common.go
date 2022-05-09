package utils

import "os"

func IsLockFilePresent() bool {

	path := "./borzoi-lock.json"

	// Check if the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}

}

func IsConfigFilePresent() bool {

	path := "./borzoi.json"

	// Check if the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}

}
