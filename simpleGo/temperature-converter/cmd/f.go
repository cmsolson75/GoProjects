/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cmsolson75/GoProjects/simpleGo/temperature-converter/conversion"
	"github.com/spf13/cobra"
)

var fCmd = &cobra.Command{
	Use:   "f",
	Short: "fahrenheit",
	Long:  `This command sets the base of the command to fahrenheit as the root.`,
	Run: func(cmd *cobra.Command, args []string) {

		if !cmd.Flags().Changed("value") || !cmd.Flags().Changed("unit") {
			fmt.Println("Please specify both the temperature and the unit. Example: --value 32.0 --unit c")
			return
		}

		// Validate the unit flag
		if T.Unit != conversion.Kelvin && T.Unit != conversion.Celsius {
			fmt.Println("Invalid Unit. Please use 'k' for Kelvin or 'c' for Celsius")
			return
		}

		switch T.Unit {
		case conversion.Kelvin:
			fmt.Printf("%.2f°F converted to Kelvin = %.2fK\n", T.Value, conversion.FahrenheitToKelvin(T.Value))
		case conversion.Celsius:
			fmt.Printf("%.2f°F converted to Celcius = %.2f°C\n", T.Value, conversion.FahrenheitToCelcius(T.Value))
		}
	},
}

func init() {
	fCmd.Flags().Float64VarP(&T.Value, "value", "v", 0.0, "Temperature value in Fahrenheit")
	fCmd.Flags().StringVarP(&T.Unit, "unit", "u", "", "Target unit for conversion: 'k' for Kelvin or 'c' for Celcius")
	rootCmd.AddCommand(fCmd)
}
