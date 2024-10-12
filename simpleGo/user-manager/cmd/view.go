/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cmsolson75/GoProjects/simpleGo/user-manager/db"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "view the csv contents",
	Long: `view the csv with eather the --all flag to see full CSV, or --email to search 
	for an email in the csv.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if viewAll {
			err := db.UserDB.ViewAll()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		} else if len(viewEmail) != 0 {
			err := db.UserDB.ViewEmail(viewEmail)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

		} else {
			fmt.Println("Please specify eather --all to see all users in db or --email to view a single email.")
		}
	},
}

var viewAll bool
var viewEmail string

func init() {
	viewCmd.Flags().BoolVarP(&viewAll, "all", "a", false, "Set to true if you want to view full db")
	viewCmd.Flags().StringVarP(&viewEmail, "email", "e", "", "Email to search in the db")

	rootCmd.AddCommand(viewCmd)
}
