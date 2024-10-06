package main

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

func checkWeightOutput(t *testing.T, got, want float64) {
	if !withinTolerance(got, want, 1e-10) {
		t.Errorf("got %.18f want %.18f", got, want)
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

func TestOzToLbs(t *testing.T) {
	got := OzToLbs(15.0)
	want := 0.9375
	checkWeightOutput(t, got, want)
}

func TestOzToG(t *testing.T) {
	got := OzToG(15.0)
	want := 425.2428
	checkWeightOutput(t, got, want)
}

func TestOzToKg(t *testing.T) {
	got := OzToKg(15.0)
	want := 0.4252428
	checkWeightOutput(t, got, want)
}

func TestOzConverter(t *testing.T) {
	got := OzConverter("lb", 15.0)
	want := 0.9375
	checkWeightOutput(t, got, want)
}

// ---- G converter
func TestGToKg(t *testing.T) {
	got := GToKg(15.0)
	want := 0.015
	checkWeightOutput(t, got, want)

}

func TestGToLbs(t *testing.T) {
	got := GToLbs(15.0)
	want := 0.0330693393277
	checkWeightOutput(t, got, want)
}

func TestGToOz(t *testing.T) {
	got := GToOz(15.0)
	exp := 0.52910942925
	checkWeightOutput(t, got, exp)
}

func TestGConverter(t *testing.T) {
	got := GConverter("oz", 15)
	want := 0.52910942925
	checkWeightOutput(t, got, want)
}

// // ---- Kg converter

func TestKgToG(t *testing.T) {
	got := KgToG(15.0)
	want := 15000.0
	checkWeightOutput(t, got, want)
}

func TestKgToOz(t *testing.T) {
	got := KgToOz(15.0)
	want := 529.10942925
	checkWeightOutput(t, got, want)
}

func TestKgToLbs(t *testing.T) {
	got := KgToLbs(15.0)
	want := 33.0693393278
	checkWeightOutput(t, got, want)
}

func TestKgConverter(t *testing.T) {
	got := KgConverter("oz", 15)
	want := 529.10942925
	checkWeightOutput(t, got, want)
}
