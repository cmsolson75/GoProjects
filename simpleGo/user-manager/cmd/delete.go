/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cmsolson75/GoProjects/simpleGo/user-manager/db"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete user given email",
	Long:  `delete user from the csv file given an email input`,
	Run: func(cmd *cobra.Command, args []string) {
		err := db.UserDB.Delete(deleteEmail)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Successfully Deleted:", deleteEmail)
	},
}

var deleteEmail string

func init() {
	deleteCmd.Flags().StringVarP(&deleteEmail, "email", "e", "", "Delete email with --email <email_to_delete@email.com>")
	rootCmd.AddCommand(deleteCmd)
}
