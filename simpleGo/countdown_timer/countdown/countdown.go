package countdown

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/cmsolson75/GoProjects/simpleGo/countdown_timer/audio"
)

const (
	minWidth        = 3
	tabWidth        = 0
	padWidth        = 0
	padChar         = '0'
	writerFlags     = tabwriter.AlignRight
	magenta         = "\033[35m"
	reset           = "\033[0m"
	clearTerminal   = "\033[H"
	moveCurserStart = "\033[2J"
)

func GetUserInput(r io.Reader, inputMessage string) (string, error) {
	reader := bufio.NewReader(r)
	fmt.Print(inputMessage)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("error reading input: " + err.Error())
	}
	input = strings.TrimSpace(input)

	return input, nil
}

func ParseUserInput(userInput string) (int, error) {
	tokens := strings.Split(userInput, ":")
	itokens := make([]int, len(tokens))

	for i, token := range tokens {
		if token == "" {
			itokens[i] = 0
		} else {
			val, err := strconv.Atoi(token)
			if err != nil {
				return 0, fmt.Errorf("invalid input: %v", err)
			}
			itokens[i] = val
		}
	}

	switch len(itokens) {
	case 1:
		return itokens[0], nil
	case 2:
		return itokens[0]*60 + itokens[1], nil
	case 3:
		return itokens[0]*3600 + itokens[1]*60 + itokens[2], nil
	default:
		return 0, errors.New("unexpected input format")
	}
}

type TimeData struct {
	hours   int
	minutes int
	seconds int
}

func TimeConverter(seconds int) TimeData {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60

	return TimeData{hours: hours, minutes: minutes, seconds: seconds}
}

func (t TimeData) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", t.hours, t.minutes, t.seconds)
}

type Writer struct {
	writer *tabwriter.Writer
	buffer *bytes.Buffer
}

func NewWriter() *Writer {
	b := &bytes.Buffer{}
	t := tabwriter.NewWriter(b, minWidth, tabWidth, padWidth, padChar, writerFlags)

	return &Writer{writer: t, buffer: b}
}

func (w *Writer) Write(t TimeData) {
	fmt.Fprintf(w.writer, "%02d:\t%02d:\t%02d \t\t", t.hours, t.minutes, t.seconds)
	w.writer.Flush()
}

func (w *Writer) Clear() {
	w.buffer.Reset()
}

func CreateColorBox(content string, boxColor string) string {
	content = strings.TrimSpace(content)
	line := strings.Repeat("─", len(content))

	return fmt.Sprintf(
		"%s╭%s╮\n%s│%s%s%s│\n%s╰%s╯\n%s",
		boxColor, line,
		boxColor, reset, content, boxColor,
		boxColor, line, reset)
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second)
}

func clearScreen(out io.Writer) {
	fmt.Fprint(out, clearTerminal)
	fmt.Fprint(out, moveCurserStart)
}

// Could Inject CreateColorBox with a interface
func Countdown(seconds int, w *Writer, out io.Writer, sleeper Sleeper, audioPlayer audio.Player) {
	clearScreen(out)
	for i := seconds; i > 0; i-- {
		timeData := TimeConverter(i)
		w.Write(timeData)
		fmt.Fprint(out, CreateColorBox(w.buffer.String(), magenta))
		sleeper.Sleep()
		clearScreen(out)
		w.Clear()
	}
	audioPlayer.Play()
}
