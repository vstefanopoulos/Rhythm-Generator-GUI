package rhythmgenerator

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

var stopPlayChan = make(chan struct{})

func play(par *par, w *widgets) {
	if *par.pattern == "" {
		return
	}
	on := makeBuffer("./wav/rim.wav")
	filler := makeBuffer("./wav/side.wav")
	off := makeBuffer("./wav/hh.wav")
	clickDownBeat := makeBuffer("./wav/clickLow.wav")
	click := makeBuffer("./wav/click.wav")

	ticker := time.NewTicker(time.Duration(60000/par.bpm) * time.Millisecond)
	var barCount int

	for {
		barCount++
		w.bar.SetText(fmt.Sprint(barCount))
		for i, char := range *par.pattern {
			select {
			case <-stopPlayChan:
				ticker.Stop()
				return
			case <-ticker.C:
				go func() {
					if w.clickCheckbox.Checked {
						switch {
						case i == 0:
							playClick(clickDownBeat)
						case w.doubletimeCheckbox.Checked && i%4 == 0:
							playClick(click)
						case i%2 == 0:
							playClick(click)
						}
					}
				}()
				go playPattern(char, w, on, filler, off)
			}
		}
	}
}

func stop() {
	stopPlayChan <- struct{}{}
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

func playPattern(char rune, w *widgets, on, filler, off *beep.Buffer) {
	switch {
	case char == 'X':
		go func() {
			snr := on.Streamer(0, on.Len())
			speaker.Play(snr)
		}()
	case char == 'x':
		go func() {
			if !w.playFillsCheckbox.Checked {
				side := filler.Streamer(0, filler.Len())
				speaker.Play(side)
			} else {
				return
			}
		}()
	case char == 'o':
		go func() {
			if w.playOffsetsCheckbox.Checked {
				hh := off.Streamer(0, off.Len())
				speaker.Play(hh)
			} else {
				return
			}
		}()
	}
}

func playClick(click *beep.Buffer) {
	streamer := click.Streamer(0, click.Len())
	speaker.Play(streamer)
}
