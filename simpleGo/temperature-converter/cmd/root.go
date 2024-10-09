/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/cmsolson75/GoProjects/simpleGo/temperature-converter/conversion"
	"github.com/spf13/cobra"
)

var T conversion.Temperature

var rootCmd = &cobra.Command{
	Use:   "temperature-converter",
	Short: "Convert between celsius, kelvin, and fahrenheit.",
	Long:  `A cli application that lets you convert the temperature of your choice.`,
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
