package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LbsToKg(weight float64) float64 {
	return weight * 0.453592
}

func LbsToG(weight float64) float64 {
	return weight * 453.59237
}

func LbsToOz(weight float64) float64 {
	return weight * 16.0
}

func LbsConverter(unit string, weight float64) float64 {
	switch unit {
	case "lb", "lbs":
		return weight
	case "g":
		output := LbsToG(weight)
		return output
	case "kg":
		output := LbsToKg(weight)
		return output
	case "oz":
		output := LbsToOz(weight)
		return output
	}
	return weight
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input a weight (float)")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	weight, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Input a unit to convert")
	unit, _ := reader.ReadString('\n')
	unit = strings.ToLower(strings.TrimSpace(unit))
	newWeight := LbsConverter(unit, weight)
	fmt.Println("New Weight: ", newWeight)

}
