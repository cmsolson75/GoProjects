package main

import "testing"

func checkOpResponce(t *testing.T, got, want float64, err error, wantErr bool) {
	if got != want || (err != nil) != wantErr {
		t.Errorf("got %f want %f, err %t err_want %t", got, want, err, wantErr)
	}
}

func TestAdd(t *testing.T) {
	got, err := Add(2, 3)
	want := 5.0
	checkOpResponce(t, got, want, err, false)
}

func TestSubtract(t *testing.T) {
	got, err := Subtract(4, 3)
	want := 1.0
	checkOpResponce(t, got, want, err, false)
}

func TestMultiply(t *testing.T) {
	t.Run("float with decimals", func(t *testing.T) {
		got, err := Multiply(3.5, 4.0)
		want := 14.0
		checkOpResponce(t, got, want, err, false)
	})
	t.Run("whole numbers", func(t *testing.T) {
		got, err := Multiply(3, 5)
		want := 15.0
		checkOpResponce(t, got, want, err, false)
	})
}

func TestDivide(t *testing.T) {
	t.Run("normal div", func(t *testing.T) {
		num, err := Divide(5, 2)
		num_want := 2.5
		checkOpResponce(t, num, num_want, err, false)
	})
	t.Run("div by zero", func(t *testing.T) {
		num, err := Divide(5, 0)
		num_want := 0.0
		checkOpResponce(t, num, num_want, err, true)
	})
}
