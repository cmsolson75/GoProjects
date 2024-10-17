package main

import "testing"

func TestParseUserInput(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "10", want: 10},
		{input: "1:00", want: 60},
		{input: "2:00", want: 120},
		{input: "02:10", want: 130},
		{input: "00:02:10", want: 130},
		{input: "01:20:10", want: 4810},
	}

	for _, tt := range tests {
		// add in error handling in the future
		got := ParseUserInput(tt.input)
		if got != tt.want {
			t.Errorf("got %d want %d", got, tt.want)
		}
	}
}
