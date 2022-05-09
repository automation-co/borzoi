/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/automation-co/borzoi/internal/lib"
	"github.com/spf13/cobra"
)

var freezeCmd = &cobra.Command{
	Use:   "freeze",
	Short: "Generates borzoi-lock.json",
	Long:  `Saves the state of the current project in a file called borzoi-lock.json.`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.Freeze()
	},
}

func init() {
	rootCmd.AddCommand(freezeCmd)
}
