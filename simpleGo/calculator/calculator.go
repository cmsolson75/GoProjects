package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func Add(a, b float64) (float64, error) {
	return a + b, nil
}

func Subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func Multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func getArgs() (float64, float64, error) {
	if flag.NArg() < 2 {
		return 0, 0, errors.New("insufficient arguments provided")
	}

	// parse a
	a, err := strconv.ParseFloat(flag.Arg(0), 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid first argument: %v", err)
	}

	// parse b
	b, err := strconv.ParseFloat(flag.Arg(1), 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid first argument: %v", err)
	}
	return a, b, nil
}

func main() {
	op := flag.String("op", "", "operation to perform: add, sub, mul, div")
	flag.Parse()

	a, b, err := getArgs()

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var result float64

	switch *op {
	case "add":
		result, err = Add(a, b)
	case "sub":
		result, err = Subtract(a, b)
	case "mul":
		result, err = Multiply(a, b)
	case "div":
		result, err = Divide(a, b)
	default:
		fmt.Println("Invalid operation. Please use -op with add, sub, mul, or div")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %2f\n", result)
}
