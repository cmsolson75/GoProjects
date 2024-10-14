package calc

import (
	"errors"
	"math"
	"strconv"
)

var ErrNegativeNumberInput = errors.New("non positive input encountered")

type CompoundInterestData struct {
	Principle    float64
	InterestRate float64
	Time         float64
}

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

func (c *CompoundInterestData) Compute() (float64, error) {
	if c.Principle <= 0 || c.InterestRate <= 0 || c.Time <= 0 {
		return 0.0, ErrNegativeNumberInput
	}

	amount := roundFloatTwoDecimalPlaces(c.Principle * math.Pow(1+c.InterestRate/100.0, c.Time))
	return amount, nil
}

func InitCompoundInterestDataString(principle, rate, time string) (*CompoundInterestData, error) {
	floatPrinciple, err := strconv.ParseFloat(principle, 64)
	if err != nil {
		return &CompoundInterestData{}, err
	}
	floatRate, err := strconv.ParseFloat(rate, 64)
	if err != nil {
		return &CompoundInterestData{}, err
	}
	floatTime, err := strconv.ParseFloat(time, 64)
	if err != nil {
		return &CompoundInterestData{}, err
	}

	return &CompoundInterestData{Principle: floatPrinciple, InterestRate: floatRate, Time: floatTime}, nil

}
