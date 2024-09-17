package rhythmgenerator

import (
	"fmt"
	"os/exec"
	"time"
)

const on = "./wav/rim.wav"
const filler = "./wav/side.wav"
const off = "./wav/hh.wav"

func playSound(sound string) {
	cmd := exec.Command("afplay", sound)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error playing sound:", err)
	}
}

func PlayPattern(pattern string, bpm int, stopChannel chan struct{}) {
	durationPerBeat := time.Duration(60000/bpm) * time.Millisecond

	for _, char := range pattern {
		select {
		case <-stopChannel:
			return
		default:
			switch char {
			case 'X':
				go playSound(on)
			case 'x':
				go playSound(filler)
			case 'o':
				go playSound(off)
			}
			time.Sleep(durationPerBeat)
		}
	}
}
