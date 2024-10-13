package main

import (
	"errors"
	"math"
)

var ErrNegativeNumberInput = errors.New("non positive input encountered")

func roundFloatTwoDecimalPlaces(f float64) float64 {
	return math.Round(f*100) / 100
}

func CompoundInterestCalculation(principle, rate, time float64) (float64, error) {
	if principle <= 0 || rate <= 0 || time <= 0 {
		return 0.0, ErrNegativeNumberInput
	}
	// round to 1e-2 for float concistancy
	amount := roundFloatTwoDecimalPlaces(principle * math.Pow(1+rate/100.0, time))
	return amount, nil
}
