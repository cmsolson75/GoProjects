/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cmsolson75/GoProjects/simpleGo/temperature-converter/conversion"
	"github.com/spf13/cobra"
)

// cCmd represents the c command
var cCmd = &cobra.Command{
	Use:   "c",
	Short: "celsius",
	Long:  `This command sets the base of the command to celsius as the root.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !cmd.Flags().Changed("value") || !cmd.Flags().Changed("unit") {
			fmt.Println("Please specify both the temperature and the unit. Example: --value 32.0 --unit f")
			return
		}

		// Validate the unit flag
		if T.Unit != conversion.Kelvin && T.Unit != conversion.Fahrenheit {
			fmt.Println("Invalid Unit. Please use 'k' for Kelvin or 'f' for Fahrenheit")
			return
		}

		switch T.Unit {
		case conversion.Kelvin:
			fmt.Printf("%.2f°C converted to Kelvin = %.2fK\n", T.Value, conversion.CelciusToKelvin(T.Value))
		case conversion.Fahrenheit:
			fmt.Printf("%.2f°C converted to Fahrenheit = %.2f°F\n", T.Value, conversion.CelciusToFahrenheit(T.Value))
		}

	},
}

func init() {
	cCmd.Flags().Float64VarP(&T.Value, "value", "v", 0.0, "Temperature value in Celsius")
	cCmd.Flags().StringVarP(&T.Unit, "unit", "u", "", "Target unit for conversion: 'k' for Kelvin or 'f' for Fahrenheit")
	rootCmd.AddCommand(cCmd)
}
