/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cmsolson75/GoProjects/simpleGo/temperature-converter/conversion"
	"github.com/spf13/cobra"
)

// kCmd represents the k command
var kCmd = &cobra.Command{
	Use:   "k",
	Short: "kelvin",
	Long:  `This command sets the base of the command to kelvin as the root.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !cmd.Flags().Changed("value") || !cmd.Flags().Changed("unit") {
			fmt.Println("Please specify both the temperature and the unit. Example: --value 32.0 --unit c")
			return
		}

		if T.Unit != conversion.Celsius && T.Unit != conversion.Fahrenheit {
			fmt.Println("Invalid Unit. Please use 'f' for Fahrenheit or 'c' for Celsius")
			return
		}

		switch T.Unit {
		case conversion.Celsius:
			fmt.Printf("%.2fK converted to Celsius = %.2f°C\n", T.Value, conversion.KelvinToCelcius(T.Value))
		case conversion.Fahrenheit:
			fmt.Printf("%.2fK converted to Fahrenheit = %.2f°F\n", T.Value, conversion.KelvinToFahrenheit(T.Value))

		}
	},
}

func init() {
	kCmd.Flags().Float64VarP(&T.Value, "value", "v", 0.0, "Temperature value in Kelvin")
	kCmd.Flags().StringVarP(&T.Unit, "unit", "u", "", "Target unit for conversion: 'f' for Fahrenheit or 'c' for Celcius")
	rootCmd.AddCommand(kCmd)
}
