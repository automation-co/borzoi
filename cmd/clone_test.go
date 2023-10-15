package cmd

import (
	"bytes"
	"os"
	"testing"
)

var (
	confFiles = []string{"borzoi.json", "borzoi-lock.json"}
)

func Test_Clone_LockFile_Present(t *testing.T) {
	// Case to test
	// Clone -> freeze -> Clone
	command := new(bytes.Buffer)
	rootCmd.SetOut(command)
	rootCmd.SetErr(command)
	// generate lock file
	rootCmd.SetArgs([]string{"generate"})
	Execute()
	// clone repo
	rootCmd.SetArgs([]string{"clone", "test.git"})
	err := rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}
	// freeze
	rootCmd.SetArgs([]string{"freeze"})
	Execute()
	rootCmd.SetArgs([]string{"clone", "test.git"})
	err = rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}

	defer func() {
		for _, v := range confFiles {
			err := os.Remove(v)
			if err != nil {
				t.Errorf("error tearing down test case :: %v", err)
			}
		}

	}()
}

func Test_Clone_LockFile_Present_Config_Missing(t *testing.T) {

	// Case to Test
	// Clone -> freeze -> Clone (but missing config)
	command := new(bytes.Buffer)
	rootCmd.SetOut(command)
	rootCmd.SetErr(command)
	rootCmd.SetArgs([]string{"generate"})
	Execute()
	rootCmd.SetArgs([]string{"clone", "test.git"})
	Execute()
	rootCmd.SetArgs([]string{"freeze"})
	Execute()
	os.Remove(confFiles[0])
	rootCmd.SetArgs([]string{"clone", "test.git"})
	Execute()

	defer func() {
		err := os.Remove(confFiles[1])
		if err != nil {
			t.Errorf("error tearing down test case :: %v", err)
		}
	}()
}
