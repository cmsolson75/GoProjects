package main

import (
	"fmt"
	"os"

	"github.com/cmsolson75/GoProjects/simpleGo/countdown_timer/audio"
	"github.com/cmsolson75/GoProjects/simpleGo/countdown_timer/countdown"
)

const (
	alarmSound = "./media/alarm.mp3"
)

func main() {
	audioInstance := audio.GetAudioPlayer()
	err := audioInstance.LoadAudio(alarmSound)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	input, err := countdown.GetUserInput(os.Stdin, "Timer Length: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	w := countdown.NewWriter()

	seconds, err := countdown.ParseUserInput(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	countdown.Countdown(seconds, w, os.Stdout, &countdown.DefaultSleeper{}, audioInstance)
	audioInstance.Quit()

}
