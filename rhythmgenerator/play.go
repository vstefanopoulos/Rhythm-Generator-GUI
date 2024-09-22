package rhythmgenerator

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

// const on = "./wav/rim.wav"
// const filler = "./wav/side.wav"
// const off = "./wav/hh.wav"

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

func makeBuffer(file string) *beep.Buffer {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/35))

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()
	return buffer
}

func playPattern(pattern string, bpm int, playFills, playOffset bool) {
	durationPerBeat := time.Duration(60000/bpm) * time.Millisecond

	on := makeBuffer("./wav/rim.wav")
	filler := makeBuffer("./wav/side.wav")
	off := makeBuffer("./wav/hh.wav")

	for _, char := range pattern {
		select {
		case <-stopPlayPatternChan:
			return
		default:
			switch char {
			case 'X':
				go func() {
					snr := on.Streamer(0, on.Len())
					speaker.Play(snr)
				}()
			case 'x':
				if playFills {
					go func() {
						side := filler.Streamer(0, filler.Len())
						speaker.Play(side)
					}()
				} else {
					time.Sleep(durationPerBeat)
					continue
				}

			case 'o':
				if playOffset {
					go func() {
						hh := off.Streamer(0, off.Len())
						speaker.Play(hh)
					}()
				} else {
					time.Sleep(durationPerBeat)
					continue
				}

			}
			time.Sleep(durationPerBeat)
		}
	}
}
