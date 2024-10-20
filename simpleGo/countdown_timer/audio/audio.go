package audio

import (
	"bytes"
	"errors"
	"io"
	"os"
	"sync"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

const (
	sr       = 44100
	channels = 2
	format   = oto.FormatSignedInt16LE
)

type Player interface {
	LoadAudio(file string) error
	Play() error
	Quit() error
}

type AudioPlayer struct {
	audioData *mp3.Decoder
	context   oto.Context
	player    oto.Player
}

func (a *AudioPlayer) LoadAudio(audioFilePath string) error {
	fileBytes, err := os.ReadFile(audioFilePath)
	if err != nil {
		return errors.New("reading file failed: " + err.Error())
	}
	fileBytesReader := bytes.NewReader(fileBytes)

	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		// Handle Errors
		return errors.New("mp3.NewDecoder failed: " + err.Error())
	}
	a.audioData = decodedMp3

	op := &oto.NewContextOptions{}

	op.SampleRate = sr
	op.ChannelCount = channels
	op.Format = format
	otoCtx, readyChan, err := oto.NewContext(op)

	if err != nil {
		return errors.New("oto.NewContex failed: " + err.Error())
	}
	// wait for hardware to be ready
	<-readyChan

	a.context = *otoCtx
	a.player = *a.context.NewPlayer(a.audioData)

	return nil
}

func (a *AudioPlayer) Play() error {
	a.player.Play()
	for a.player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	// Reset play head
	_, err := a.player.Seek(0, io.SeekStart)
	if err != nil {
		return errors.New("player.Seek failed: " + err.Error())
	}

	return nil

}

func (a *AudioPlayer) Quit() error {
	err := a.player.Close()
	if err != nil {
		return errors.New("error closing file: " + err.Error())
	}
	return nil
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
