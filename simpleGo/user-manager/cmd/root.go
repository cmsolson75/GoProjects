/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "user-manager",
	Short: "manage users in a CSV db",
	Long: `Add, Remove, and View users in a CSV database,
	to add users use the add command with --email, to remove users
	use the remove command with --email, To view users use view with eather
	--all to view all, or --email <youremail@email.com>`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
