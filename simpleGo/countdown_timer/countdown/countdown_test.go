package countdown

import (
	"bytes"
	"strings"
	"testing"
)

func TestGetUserInput(t *testing.T) {
	reader := strings.NewReader("10\n")
	got, _ := GetUserInput(reader, "")
	want := "10"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}

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
		{input: "::10", want: 10},
	}

	for _, tt := range tests {
		// add in error handling in the future
		got, _ := ParseUserInput(tt.input)
		if got != tt.want {
			t.Errorf("got %d want %d", got, tt.want)
		}
	}
}

func TestWriterWrite(t *testing.T) {
	w := NewWriter()

	timeData := TimeData{hours: 10, minutes: 11, seconds: 9}

	w.Write(timeData)
	got := w.buffer.String()
	want := "10:11:09 "

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestTimeConverter(t *testing.T) {
	tests := []struct {
		input   int
		hours   int
		minutes int
		seconds int
	}{
		{input: 300, hours: 0, minutes: 5, seconds: 0},
		{input: 210, hours: 0, minutes: 3, seconds: 30},
		{input: 6000, hours: 1, minutes: 40, seconds: 0},
	}

	for _, tt := range tests {
		d := TimeConverter(tt.input)

		assertTimeUnit(t, "hours", d.hours, tt.hours)
		assertTimeUnit(t, "minutes", d.minutes, tt.minutes)
		assertTimeUnit(t, "seconds", d.seconds, tt.seconds)
	}
}

func TestCreateColorBox(t *testing.T) {
	input := "abc"
	// Color Escape Codes
	want := "\x1b[35m╭───╮\n\x1b[35m│\x1b[0mabc\x1b[35m│\n\x1b[35m╰───╯\n\x1b[0m"
	got := CreateColorBox(input, "\033[35m")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

type DummySleeper struct{}

func (d *DummySleeper) Sleep() {}

type DummyAudioPlayer struct{}

func (d *DummyAudioPlayer) LoadAudio(file string) error { return nil }
func (d *DummyAudioPlayer) Play() error                 { return nil }
func (d *DummyAudioPlayer) Quit() error                 { return nil }

// Bad test: not testing behavior
func TestCountdown(t *testing.T) {
	w := NewWriter()
	buf := bytes.Buffer{}
	Countdown(2, w, &buf, &DummySleeper{}, &DummyAudioPlayer{})
	want := "\x1b[H\x1b[2J\x1b[35m╭────────╮\n\x1b[35m│\x1b[0m00:00:02\x1b[35m│\n\x1b[35m╰────────╯\n\x1b[0m\x1b[H\x1b[2J\x1b[35m╭────────╮\n\x1b[35m│\x1b[0m00:00:01\x1b[35m│\n\x1b[35m╰────────╯\n\x1b[0m\x1b[H\x1b[2J"
	if buf.String() != want {
		t.Errorf("got %q want %q", buf.String(), want)
	}
}

func assertTimeUnit(t testing.TB, unit string, valueGot, valueWant int) {
	t.Helper()
	if valueGot != valueWant {
		t.Errorf("%s incorrect: got %d want %d", unit, valueGot, valueWant)
	}
}
