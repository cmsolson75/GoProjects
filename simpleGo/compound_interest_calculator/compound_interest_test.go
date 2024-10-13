package main

import (
	"testing"
)

func TestCompoundInterestCalculation(t *testing.T) {
	t.Run("valid inputs", func(t *testing.T) {
		p := 1000.0
		r := 10.0
		time := 10.0
		got, _ := CompoundInterestCalculation(p, r, time)
		want := 2593.74 // need to round output to 2 decimals
		assertFloat(t, got, want)
	})

	// handle before: but have the func handle if not handled before
	t.Run("negative value", func(t *testing.T) {
		p := 100.0
		r := -10.0
		time := 10.0
		_, err := CompoundInterestCalculation(p, r, time)
		want := ErrNegativeNumberInput
		if err == nil {
			t.Fatal("expected error, got no error")
		}

		if err != want {
			t.Errorf("got %s want %s", err, want)
		}
	})
}

func assertFloat(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
