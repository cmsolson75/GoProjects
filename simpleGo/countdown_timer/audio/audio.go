package audio

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

type Player interface {
	LoadAudio(file string)
	Play()
	Quit()
}

type AudioPlayer struct {
	audioData *mp3.Decoder
	context   oto.Context
	player    oto.Player
}

func (a *AudioPlayer) LoadAudio(audioFilePath string) {
	fileBytes, err := os.ReadFile(audioFilePath)
	if err != nil {
		panic("reading file failed: " + err.Error())
	}
	fileBytesReader := bytes.NewReader(fileBytes)

	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		// Handle Errors
		panic("mp3.NewDecoder failed: " + err.Error())
	}
	a.audioData = decodedMp3

	op := &oto.NewContextOptions{}

	op.SampleRate = 44100
	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE
	otoCtx, readyChan, err := oto.NewContext(op)

	if err != nil {
		panic("oto.NewContex failed: " + err.Error())
	}
	// wait for hardware to be ready
	<-readyChan

	a.context = *otoCtx
	a.player = *a.context.NewPlayer(a.audioData)
}

func (a *AudioPlayer) Play() {
	a.player.Play()
	for a.player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	// Reset play head
	_, err := a.player.Seek(0, io.SeekStart)
	if err != nil {
		panic("player.Seek failed: " + err.Error())
	}

}

func (a *AudioPlayer) Quit() {
	err := a.player.Close()
	if err != nil {
		panic(err)
	}
}

// Singleton to avoid erros with oto.Context
var instance *AudioPlayer
var once sync.Once

func GetAudioPlayer() *AudioPlayer {
	once.Do(func() {
		instance = &AudioPlayer{}
	})
	return instance
}
