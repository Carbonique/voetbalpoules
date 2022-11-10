/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Overridden at build time
var Version string

// voorspellingenCmd represents the voorspellingen command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print client version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %v\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
