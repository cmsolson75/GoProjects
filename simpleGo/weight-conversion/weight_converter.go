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
	default:
		return weight
	}
}

func OzToLbs(weight float64) float64 {
	return weight * 0.0625
}

func OzToG(weight float64) float64 {
	return weight * 28.34952
}

func OzToKg(weight float64) float64 {
	return weight * 0.02834952
}

func OzConverter(unit string, weight float64) float64 {
	switch unit {
	case "lb", "lbs":
		return OzToLbs(weight)
	case "g":
		return OzToG(weight)
	case "kg":
		return OzToKg(weight)
	default:
		return weight
	}
}

func GToKg(weight float64) float64 {
	return weight / 1000
}

func GToLbs(weight float64) float64 {
	return weight * 0.00220462262185
}

func GToOz(weight float64) float64 {
	return weight * 0.03527396195
}

func GConverter(unit string, weight float64) float64 {
	switch unit {
	case "lb", "lbs":
		return GToLbs(weight)
	case "oz":
		return GToOz(weight)
	case "kg":
		return GToKg(weight)
	default:
		return weight
	}
}

func KgToG(weight float64) float64 {
	return weight * 1000
}

func KgToOz(weight float64) float64 {
	return weight * 35.27396195
}

func KgToLbs(weight float64) float64 {
	return weight * 2.20462262185
}

func KgConverter(unit string, weight float64) float64 {
	switch unit {
	case "lb", "lbs":
		return KgToLbs(weight)
	case "oz":
		return KgToOz(weight)
	case "g":
		return KgToG(weight)
	default:
		return weight
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input a weight (float) and a unit (lbs, g, kg, oz)")
	input, _ := reader.ReadString('\n')
	line := strings.TrimSpace(input)
	inputs := strings.Fields(line)
	if len(inputs) < 2 {
		fmt.Println("Not enough inputs exiting...")
		os.Exit(1)
	}
	weight, err := strconv.ParseFloat(inputs[0], 64)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	inputUnit := strings.ToLower(inputs[1])

	fmt.Println("Input a unit to convert")
	unit, _ := reader.ReadString('\n')
	unit = strings.ToLower(strings.TrimSpace(unit))
	var newWeight float64

	switch inputUnit {
	case "lb", "lbs", "pounds", "pound":
		newWeight = LbsConverter(unit, weight)
	case "oz", "ounce":
		newWeight = OzConverter(unit, weight)
	case "kg", "kgs", "kilogram", "kilograms":
		newWeight = KgConverter(unit, weight)
	case "g", "gs", "gram", "grams":
		newWeight = GConverter(unit, weight)
	default:
		newWeight = weight
	}

	fmt.Println("New Weight: ", newWeight)

}
