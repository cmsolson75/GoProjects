package main

import "testing"

func checkOpResponce(t *testing.T, num_got, num_want float64, err, err_want bool) {
	if num_got != num_want || err != err_want {
		t.Errorf("got %f want %f, err %t err_want %t", num_got, num_want, err, err_want)
	}
}

func TestAdd(t *testing.T) {
	got, err := Add(2, 3)
	want, err_want := 5.0, false
	checkOpResponce(t, got, want, err, err_want)
}

func TestSubtract(t *testing.T) {
	got, err := Subtract(4, 3)
	want, err_want := 1.0, false
	checkOpResponce(t, got, want, err, err_want)
}

func TestMultiply(t *testing.T) {
	t.Run("float with decimals", func(t *testing.T) {
		got, err := Multiply(3.5, 4.0)
		want, err_want := 14.0, false
		checkOpResponce(t, got, want, err, err_want)
	})
	t.Run("whole numbers", func(t *testing.T) {
		got, err := Multiply(3, 5)
		want, err_want := 15.0, false
		checkOpResponce(t, got, want, err, err_want)
	})
}

func TestDivide(t *testing.T) {
	t.Run("normal div", func(t *testing.T) {
		num, err := Divide(5, 2)
		num_want, err_want := 2.5, false
		checkOpResponce(t, num, num_want, err, err_want)
	})
	t.Run("div by zero", func(t *testing.T) {
		num, err := Divide(5, 0)
		num_want, err_want := 0.0, true
		checkOpResponce(t, num, num_want, err, err_want)
	})
}
