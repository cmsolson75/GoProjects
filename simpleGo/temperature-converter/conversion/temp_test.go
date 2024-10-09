package conversion

import (
	"math"
	"testing"
)

func withinTolerance(a, b, e float64) bool {
	if a == b {
		return true
	}

	d := math.Abs(a - b)

	if b == 0 {
		return d < e
	}

	return (d / math.Abs(b)) < e
}

func checkTempOutput(t *testing.T, got, want float64) {
	if !withinTolerance(got, want, 1e-10) {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestFahrenheitToCelcius(t *testing.T) {
	got := FahrenheitToCelcius(15.0)
	want := -9.4444444444
	checkTempOutput(t, got, want)

}

func TestCelciusToFahrenheit(t *testing.T) {
	got := CelciusToFahrenheit(15.0)
	want := 59.0
	checkTempOutput(t, got, want)
}

func TestKelvinToFahrenheit(t *testing.T) {
	got := KelvinToFahrenheit(15.0)
	want := -432.67
	checkTempOutput(t, got, want)
}

func TestFahrenheitToKelvin(t *testing.T) {
	got := FahrenheitToKelvin(15.0)
	want := 263.7055555556
	checkTempOutput(t, got, want)
}

func TestCelciusToKelvin(t *testing.T) {
	got := CelciusToKelvin(15.0)
	want := 288.15
	checkTempOutput(t, got, want)
}

func TestKelvinToCelcius(t *testing.T) {
	got := KelvinToCelcius(15.0)
	want := -258.15
	checkTempOutput(t, got, want)
}
