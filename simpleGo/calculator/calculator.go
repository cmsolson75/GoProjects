package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

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

func getArgs() (float64, float64) {
	// check flags
	if flag.Arg(0) == "" || flag.Arg(1) == "" {
		fmt.Println("Please provide arguments")
		os.Exit(0)
	}

	// parse a
	a, err := strconv.ParseFloat(flag.Arg(0), 64)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}

	// parse b
	b, err := strconv.ParseFloat(flag.Arg(1), 64)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
	return a, b
}

func main() {
	useAdd := flag.Bool("add", false, "use add command")
	useSub := flag.Bool("sub", false, "use subtract command")
	useMul := flag.Bool("mul", false, "use multiply command")
	useDiv := flag.Bool("div", false, "use divide command")
	flag.Parse()
	a, b := getArgs()

	if *useAdd {
		output, _ := Add(a, b)
		fmt.Printf("%.2f + %.2f = %.2f\n", a, b, output)
	} else if *useSub {
		output, _ := Subtract(a, b)
		fmt.Printf("%.2f - %.2f = %.2f\n", a, b, output)
	} else if *useMul {
		output, _ := Multiply(a, b)
		fmt.Printf("%.2f * %.2f = %.2f\n", a, b, output)
	} else if *useDiv {
		output, err := Divide(a, b)
		if err {
			fmt.Println("Cant divide by 0")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", a, b, output)
		}
	}
}
