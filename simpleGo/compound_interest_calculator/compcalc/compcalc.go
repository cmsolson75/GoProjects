package compcalc

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"unicode"

	"github.com/cmsolson75/GoProjects/simpleGo/compound_interest_calculator/state"
)

var ErrNegativeNumberInput = errors.New("non positive input encountered")
var ErrNonFloatNumberInput = errors.New("input not valid number")

type CompoundInterestData struct {
	Principle    float64
	InterestRate float64
	Time         float64
}

func roundFloatTwoDecimalPlaces(f float64) float64 {
	return math.Round(f*100) / 100
}

func (c *CompoundInterestData) Compute() (float64, error) {
	if c.Principle <= 0 || c.InterestRate <= 0 || c.Time <= 0 {
		return 0.0, ErrNegativeNumberInput
	}

	amount := roundFloatTwoDecimalPlaces(c.Principle * math.Pow(1+c.InterestRate/100.0, c.Time))
	return amount, nil
}

func (c *CompoundInterestData) ComputeWithTime(t int) (float64, error) {
	if t <= 0 {
		return 0.0, ErrNegativeNumberInput
	}
	amount := roundFloatTwoDecimalPlaces(c.Principle * math.Pow(1+c.InterestRate/100.0, float64(t)))
	return amount, nil
}

func IsValidFloat(s string) bool {

	// Handle if the input is only one number
	if len(s) == 1 {
		return unicode.IsDigit(rune(s[0]))
	}

	for i, r := range s {
		// Allow '-' if in first position
		if i == 0 {
			if !unicode.IsDigit(r) && r != '-' {
				fmt.Println("Fail in == 0")
				return false
			}
		} else {
			if !unicode.IsDigit(r) && r != '.' {
				return false
			}
		}
	}
	return true
}

func InitCompoundInterestDataString(state state.AppState) (*CompoundInterestData, error) {
	if !IsValidFloat(state.InitialInvestment) || !IsValidFloat(state.InterestRate) || !IsValidFloat(state.TimeYears) {
		return &CompoundInterestData{}, ErrNonFloatNumberInput
	}

	floatPrinciple, err := strconv.ParseFloat(state.InitialInvestment, 64)
	if err != nil {
		return &CompoundInterestData{}, err
	}
	floatRate, err := strconv.ParseFloat(state.InterestRate, 64)
	if err != nil {
		return &CompoundInterestData{}, err
	}
	floatTime, err := strconv.ParseFloat(state.TimeYears, 64)
	if err != nil {
		return &CompoundInterestData{}, err
	}

	return &CompoundInterestData{Principle: floatPrinciple, InterestRate: floatRate, Time: floatTime}, nil

}
