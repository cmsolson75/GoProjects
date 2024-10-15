package state

import "errors"

type AppState struct {
	InitialInvestment string
	InterestRate      string
	TimeYears         string
}

var ErrEmptyInput = errors.New("empty input field detected, please fill in all fields")

func (s AppState) CheckEmptyInput() error {
	if s.InitialInvestment == "" || s.InterestRate == "" || s.TimeYears == "" {
		return ErrEmptyInput
	}
	return nil
}
