package main

import (
	"bufio"
	"math"
	"reflect"
	"strings"
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

func TestWeightFactory(t *testing.T) {
	tests := []struct {
		unit   string
		weight float64
		want   Weight
		err    bool
	}{
		{"pound", 150.0, Weight{"pound", 150.0, map[string]float64{"ounce": 16.0, "gram": 453.59237, "kilogram": 0.453592, "pound": 1.0}}, false},
		{"ounce", 100.0, Weight{"ounce", 100.0, map[string]float64{"ounce": 1.0, "gram": 28.34952, "kilogram": 0.02834952, "pound": 0.0625}}, false},
		{"invalid", 100.0, Weight{}, true},
	}

	for _, tt := range tests {
		got, err := WeightFactory(tt.unit, tt.weight)
		if (err != nil) != tt.err {
			t.Errorf("WeightFactory() error = %v, wantErr %v", err, tt.err)
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("WeightFactory() = %v, want %v", got, tt.want)
		}
	}
}

func TestConvertWeight(t *testing.T) {
	w, err := WeightFactory("pound", 150.0)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	got, err := w.ConvertWeight("kilogram")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	want := 68.0388
	checkWeightOutput(t, got, want)
}

func TestInvalidConversion(t *testing.T) {
	w, err := WeightFactory("pound", 150.0)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	_, err = w.ConvertWeight("stone")
	if err == nil {
		t.Errorf("expected error, got none")
	}
}

func TestGetValidUnit(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"lbs", "pound", false},
		{"ounces", "ounce", false},
		{"grams", "gram", false},
		{"stone", "", true},
	}

	for _, tt := range tests {
		got, err := GetValidUnit(tt.input)
		if (err != nil) != tt.err {
			t.Errorf("GetValidUnit(%v) error = %v, wantErr %v", tt.input, err, tt.err)
		}
		if got != tt.expected {
			t.Errorf("GetValidUnit(%v) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}

func TestGetWeightInput(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("150 lbs\n"))
	unit, weight, err := GetWeightInput(reader)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if unit != "pound" {
		t.Errorf("got %v, want %v", unit, "pound")
	}
	if weight != 150.0 {
		t.Errorf("got %v, want %v", weight, 150.0)
	}
}

func TestGetUnitConversionInput(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("kg\n"))
	unit, err := GetUnitConversionInput(reader)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if unit != "kilogram" {
		t.Errorf("got %v, want %v", unit, "kilogram")
	}
}
