package compcalc

import (
	"testing"
)

func TestCompoundInterestCalculation(t *testing.T) {
	t.Run("valid inputs", func(t *testing.T) {
		cid := CompoundInterestData{Principle: 1000.0, InterestRate: 10.0, Time: 10.0}
		got, err := cid.Compute()
		if err != nil {
			t.Error("error detected when not expected")
		}
		want := 2593.74 // need to round output to 2 decimals
		assertFloat(t, got, want)
	})

	t.Run("negative value", func(t *testing.T) {
		cid := CompoundInterestData{Principle: 1000.0, InterestRate: -10.0, Time: 10.0}
		_, err := cid.Compute()
		want := ErrNegativeNumberInput
		if err == nil {
			t.Fatal("expected error, got no error")
		}

		if err != want {
			t.Errorf("got %s want %s", err, want)
		}
	})
}

func TestIsValidFloat(t *testing.T) {
	t.Run("valid int string input", func(t *testing.T) {
		got := IsValidFloat("10")
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("valid float string input", func(t *testing.T) {
		got := IsValidFloat("12.1")
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("invalid float string input", func(t *testing.T) {
		got := IsValidFloat("12.1b")
		want := false
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

func assertFloat(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
