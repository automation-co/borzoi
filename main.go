package main

import (
	"github.com/automation-co/borzoi/cmd"
	"github.com/automation-co/borzoi/internal/config"
)

func init() {

	// Initialize the config
	config.InitializeConfig()
}

func main() {
	cmd.Execute()
}
