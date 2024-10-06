package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Base Weight for processing
type Weight struct {
	Unit              string
	Weight            float64
	ConversionFactors map[string]float64
}

// Create weight varients
func WeightFactory(unit string, weight float64) (Weight, error) {
	conversionData := map[string]map[string]float64{
		"pound":    {"ounce": 16.0, "gram": 453.59237, "kilogram": 0.453592, "pound": 1.0},
		"ounce":    {"ounce": 1.0, "gram": 28.34952, "kilogram": 0.02834952, "pound": 0.0625},
		"gram":     {"ounce": 0.03527396195, "gram": 1.0, "kilogram": 0.001, "pound": 0.00220462262185},
		"kilogram": {"ounce": 35.27396195, "gram": 1000.0, "kilogram": 1.0, "pound": 2.20462262185},
	}

	factors, exists := conversionData[unit]
	if !exists {
		return Weight{}, errors.New("unsupported unit")
	}
	return Weight{Unit: unit, Weight: weight, ConversionFactors: factors}, nil

}

func (w *Weight) ConvertWeight(unit string) (float64, error) {
	factor, exists := w.ConversionFactors[unit]
	if !exists {
		return 0.0, errors.New("unsupported conversion unit")
	}
	return factor * w.Weight, nil
}

// normalize units
func GetValidUnit(unit string) (string, error) {
	unitMappings := map[string]string{
		"lb":        "pound",
		"lbs":       "pound",
		"pounds":    "pound",
		"pound":     "pound",
		"oz":        "ounce",
		"ounces":    "ounce",
		"ounce":     "ounce",
		"kg":        "kilogram",
		"kgs":       "kilogram",
		"kilogram":  "kilogram",
		"kilograms": "kilogram",
		"g":         "gram",
		"grams":     "gram",
		"gram":      "gram",
		"gs":        "gram",
	}

	normalizedUnit, exists := unitMappings[unit]
	if !exists {
		return "", fmt.Errorf("invalid unit: %s", unit)
	}

	return normalizedUnit, nil
}

// get weight and unit for starting weight
func GetWeightInput(reader *bufio.Reader) (string, float64, error) {
	lineInput, err := reader.ReadString('\n')
	if err != nil {
		return "", 0.0, err
	}
	// clean input
	lineInput = strings.TrimSpace(lineInput)

	// extract sub strings
	userInputs := strings.Fields(lineInput)
	// check if there are enough arguments
	if len(userInputs) < 2 {
		return "", 0.0, errors.New("not enough arguments")
	}

	weightArg, unitArg := userInputs[0], userInputs[1]

	// convert to float64
	weight, err := strconv.ParseFloat(weightArg, 64)
	if err != nil {
		return "", 0.0, err
	}
	// normalize for str matching
	unit := strings.ToLower(unitArg)

	// check valid
	unit, err = GetValidUnit(unit)
	if err != nil {
		return "", 0.0, err
	}

	return unit, weight, nil

}

// get valid input unit
func GetUnitConversionInput(reader *bufio.Reader) (string, error) {
	fmt.Println("Input a unit to convert")
	unit, _ := reader.ReadString('\n')
	unit = strings.ToLower(strings.TrimSpace(unit))
	unit, err := GetValidUnit(unit)
	if err != nil {
		return "", err
	}
	return unit, nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input a weight (float) and a unit (lbs, g, kg, oz)")
	inputUnit, weight, err := GetWeightInput(reader)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	wf, err := WeightFactory(inputUnit, weight)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	unit, err := GetUnitConversionInput(reader)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	convertedWeight, err := wf.ConvertWeight(unit)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Converted Weight: ", convertedWeight)

}
