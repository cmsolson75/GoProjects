package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

// Need errors
// Need input sanitization: non numbers, invalid numbers
// HANDLE THE ERRORS
// The extraction is based on length: so I could just extract and opend to a list and then
// make it work

func GetUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("TIME: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)

	return input
}

func ParseUserInput(userInput string) int {
	delimiters := ":"
	tokens := strings.Split(userInput, delimiters)

	var itokens []int
	for _, t := range tokens {
		fmt.Println(t)
		i, err := strconv.Atoi(t)
		if err != nil {
			i = 0
		}
		itokens = append(itokens, i)
	}
	var output int
	switch len(itokens) {
	case 1:
		output = itokens[0]
	case 2:
		output = (itokens[0] * 60) + itokens[1]
	case 3:
		output = (itokens[0] * 3600) + (itokens[1] * 60) + itokens[2]
	}

	return output
}

func Countdown(seconds int) {
	// use DI in the future for the os.Stdout
	// Might need to write the output to a different variable for formatting
	buf := bytes.Buffer{}
	w := tabwriter.NewWriter(&buf, 3, 0, 0, '0', tabwriter.AlignRight)
	fmt.Print("\033[H\033[2J")
	for i := seconds; i > 0; i-- {
		// This is a format call for getting the relevant time objects
		h := (i / 3600)
		m := (i % 3600) / 60
		s := (i % 3600) % 60
		// This printout is anoying
		// Refactor it into a different place
		fmt.Fprintf(w, "%s:\t%s:\t%s \t\t", strconv.Itoa(h), strconv.Itoa(m), strconv.Itoa(s))
		w.Flush()
		fmt.Printf("\u001b[31mCOUNTDOWN TIMER\n  ---------- \n | %s|\n  ---------- ", buf.String())
		time.Sleep(time.Second)
		fmt.Print("\033[H\033[2J")
		buf.Reset()
	}
	fmt.Println("\u001b[31mDone!")
}

func main() {
	input := GetUserInput()

	seconds := ParseUserInput(input)
	Countdown(seconds)
}
