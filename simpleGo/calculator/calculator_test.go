package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(4, 3)
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestMultiply(t *testing.T) {
	t.Run("float with decimals", func(t *testing.T) {
		got := Multiply(3.5, 4.0)
		want := 14.0
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
	t.Run("whole numbers", func(t *testing.T) {
		got := Multiply(3, 5)
		want := 15.0
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}

func TestDivide(t *testing.T) {
	got := Divide(5, 2)
	want := 2.5
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
