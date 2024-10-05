package main

import "testing"

func checkWeightOutput(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestLbsToKg(t *testing.T) {
	got := LbsToKg(150.0)
	want := 68.0388
	checkWeightOutput(t, got, want)
}

func TestLbsToG(t *testing.T) {
	got := LbsToG(150.0)
	want := 68038.8555
	checkWeightOutput(t, got, want)
}

func TestLbsToOz(t *testing.T) {
	got := LbsToOz(150.0)
	want := 2400.0
	checkWeightOutput(t, got, want)
}

func TestLbsConverter(t *testing.T) {
	got := LbsConverter("oz", 10)
	want := 160.0
	checkWeightOutput(t, got, want)
}
