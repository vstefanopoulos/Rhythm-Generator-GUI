// TODO: Fix latency when re instaciating ticker
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
var changeBpmChan = make(chan struct{})

func play(p *Parameters, w *Widgets) {
	if *p.pattern == "" {
		return
	}

	on := makeBuffer("./wav/rim.wav")
	filler := makeBuffer("./wav/side.wav")
	off := makeBuffer("./wav/hh.wav")
	clickDownBeat := makeBuffer("./wav/clickLow.wav")
	click := makeBuffer("./wav/click.wav")

	var bpm int
	bpm = newBpm(w, p.bpm)
	ticker := time.NewTicker(time.Duration(60000/bpm) * time.Millisecond)

	var barCount int
	isPlaying = true
	for {
		barCount++
		w.barLabel.SetText(fmt.Sprint("Bar: ", barCount))
		select {
		case <-changeBpmChan:
			bpm = newBpm(w, p.bpm)
			select {
			case <-ticker.C:
				ticker = time.NewTicker(time.Duration(60000/bpm) * time.Millisecond)
			}

		default:
			for i, char := range *p.pattern {
				select {
				case <-stopPlayChan:
					ticker.Stop()
					return
				case <-ticker.C:
					go func() {
						if w.clickCheck.Checked {
							switch {
							case i == 0 && w.accentDownbeatCheck.Checked:
								playClick(clickDownBeat)
							case w.doubletimeCheck.Checked && i%4 == 0:
								playClick(click)
							case !w.doubletimeCheck.Checked && i%2 == 0:
								playClick(click)
							}
						}
					}()
					go playPattern(char, w, on, filler, off)
				}
			}
		}
	}
}

func newBpm(w *Widgets, bpm int) int {
	if w.doubletimeCheck.Checked {
		return bpm * 2
	}
	return bpm
}

func stop() {
	stopPlayChan <- struct{}{}
	isPlaying = false
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

func playPattern(char rune, w *Widgets, on, filler, off *beep.Buffer) {
	switch {
	case char == 'X':
		go func() {
			snr := on.Streamer(0, on.Len())
			speaker.Play(snr)
		}()
	case char == 'x':
		go func() {
			if !w.muteFillsCheck.Checked {
				side := filler.Streamer(0, filler.Len())
				speaker.Play(side)
			} else {
				return
			}
		}()
	case char == 'o':
		go func() {
			if !w.muteOffsetsCheck.Checked {
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
