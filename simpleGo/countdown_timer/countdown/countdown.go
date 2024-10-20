package countdown

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/cmsolson75/GoProjects/simpleGo/countdown_timer/audio"
)

const (
	minWidth    = 3
	tabWidth    = 0
	padWidth    = 0
	padChar     = '0'
	writerFlags = tabwriter.AlignRight
)

func GetUserInput(r io.Reader, inputMessage string) string {
	reader := bufio.NewReader(r)
	fmt.Print(inputMessage)

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
		i, err := strconv.Atoi(t)
		// An error isnt always bad
		// ::10, should be valid as 10 seconds
		// issue: Can have silent errors
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

type TimeData struct {
	hours   string
	minutes string
	seconds string
}

// tested
func TimeConverter(s int) TimeData {
	hours := strconv.Itoa(s / 3600)
	minutes := strconv.Itoa((s % 3600) / 60)
	seconds := strconv.Itoa((s % 3600) % 60)

	return TimeData{hours: hours, minutes: minutes, seconds: seconds}
}

type Buffer interface {
	Clear()
}

type TimeBuffer struct {
	buffer bytes.Buffer
}

func (c *TimeBuffer) Clear() {
	c.buffer.Reset()
}

func (c TimeBuffer) String() string {
	b := c.buffer
	return b.String()
}

// No need for a custom buffer like I originally thought
type Writer struct {
	writer *tabwriter.Writer
	buffer *TimeBuffer
}

func NewWriter() *Writer {
	b := TimeBuffer{}
	t := tabwriter.NewWriter(&b.buffer, minWidth, tabWidth, padWidth, padChar, writerFlags)

	return &Writer{writer: t, buffer: &b}
}

func (u *Writer) Write(t TimeData) {
	fmt.Fprintf(u.writer, "%s:\t%s:\t%s \t\t", t.hours, t.minutes, t.seconds)
	u.writer.Flush()
}

const (
	magenta = "\033[35m"
	reset   = "\033[0m"
)

func CreateColorBox(content string, boxColor string) string {
	content = strings.TrimSpace(content)

	var boxOutput string

	line := strings.Repeat("─", len(content))
	boxOutput += fmt.Sprintf("%s╭%s╮\n", boxColor, line)
	boxOutput += fmt.Sprintf("%s│%s%s%s│\n", boxColor, reset, content, boxColor)
	boxOutput += fmt.Sprintf("%s╰%s╯\n%s", boxColor, line, reset)
	return boxOutput
}

const (
	clearTerminal   = "\033[H"
	moveCurserStart = "\033[2J"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second)
}

// Could Inject CreateColorBox with a interface
func Countdown(seconds int, w *Writer, out io.Writer, sleeper Sleeper, audioPlayer audio.Player) {
	fmt.Fprint(out, clearTerminal)
	fmt.Fprint(out, moveCurserStart)
	for i := seconds; i > 0; i-- {
		timeData := TimeConverter(i)
		w.Write(timeData)
		fmt.Fprint(out, CreateColorBox(w.buffer.String(), magenta))
		sleeper.Sleep()
		fmt.Fprint(out, clearTerminal)
		fmt.Fprint(out, moveCurserStart)
		w.buffer.Clear()
	}
	audioPlayer.Play()
}
