package rhythmgenerator

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)

const on = "./wav/rim.wav"
const filler = "./wav/side.wav"
const off = "./wav/hh.wav"

var stopPlayChan = make(chan struct{})
var stopPlayPatternChan = make(chan struct{})

func play(pattern string, bpm int, w *widgets) {
	if pattern == "" {
		return
	}

	var barCount int
	var wg sync.WaitGroup
	for {
		select {
		case <-stopPlayChan:
			return
		default:
			barCount++

			wg.Add(2)

			go func() {
				defer wg.Done()
				w.bar.SetText(fmt.Sprintf("Bar: %v", barCount))
			}()

			go func() {
				defer wg.Done()
				playPattern(pattern, bpm, w.playFillsCheckbox.Checked, w.playOffsetsCheckbox.Checked)
			}()

			wg.Wait()
		}
	}

}

func stop() {
	go func() {
		stopPlayChan <- struct{}{}
	}()
	go func() {
		stopPlayPatternChan <- struct{}{}
	}()
}

func playSound(sound string) {
	cmd := exec.Command("afplay", sound)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error playing sound:", err)
	}
}

func playPattern(pattern string, bpm int, playFills, playOffset bool) {
	durationPerBeat := time.Duration(60000/bpm) * time.Millisecond

	for _, char := range pattern {
		select {
		case <-stopPlayPatternChan:
			return
		default:
			switch char {
			case 'X':
				go playSound(on)
			case 'x':
				if playFills {
					go playSound(filler)
				} else {
					time.Sleep(durationPerBeat)
					continue
				}

			case 'o':
				if playOffset {
					go playSound(off)
				} else {
					time.Sleep(durationPerBeat)
					continue
				}

			}
			time.Sleep(durationPerBeat)
		}
	}
}
