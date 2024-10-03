package main

func Add(a, b float64) (float64, bool) {
	return a + b, false
}

func Subtract(a, b float64) (float64, bool) {
	return a - b, false
}

func Multiply(a, b float64) (float64, bool) {
	return a * b, false
}

func Divide(a, b float64) (float64, bool) {
	if b != 0 {
		result := a / b
		return result, false
	} else {
		result := 0.0
		return result, true
	}
}
