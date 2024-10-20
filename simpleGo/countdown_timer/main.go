package main

import (
	"os"

	"github.com/cmsolson75/GoProjects/simpleGo/countdown_timer/audio"
	"github.com/cmsolson75/GoProjects/simpleGo/countdown_timer/countdown"
)

const (
	alarmSound = "./media/alarm.mp3"
)

func main() {
	audioInstance := audio.GetAudioPlayer()
	audioInstance.LoadAudio(alarmSound)
	input := countdown.GetUserInput(os.Stdin, "Timer Length: ")
	w := countdown.NewWriter()

	seconds := countdown.ParseUserInput(input)
	countdown.Countdown(seconds, w, os.Stdout, &countdown.DefaultSleeper{}, audioInstance)
	audioInstance.Quit()

}
