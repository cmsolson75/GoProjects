/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cmsolson75/GoProjects/simpleGo/user-manager/db"
	"github.com/cmsolson75/GoProjects/simpleGo/user-manager/slicer"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add user to the csv",
	Long:  `convert an email into username, domain and stores it in a csv file`,
	Run: func(cmd *cobra.Command, args []string) {
		userAccount, err := slicer.EmailSlicer(Email)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		err = db.UserDB.Add([]string{userAccount.Username, userAccount.Domain})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		err = db.UserDB.Write()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Success adding %s to db.\n", Email)

	},
}

var Email string

func init() {
	addCmd.Flags().StringVarP(&Email, "email", "e", "", "Email to add to DB")
	rootCmd.AddCommand(addCmd)
}
