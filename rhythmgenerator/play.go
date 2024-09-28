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
var stopClockChan = make(chan struct{})
var changeBpmChan = make(chan struct{})

func play(p *Parameters, w *Widgets, buf *Buffer) {
	if *p.pattern == "" {
		return
	}
	p.beat = newBeat(w, p.bpm)
	click := make(chan struct{})
	go clock(p, &click)
	var barCount int
	p.isPlaying = true
	for {
		barCount++
		w.barLabel.Text = (fmt.Sprint("Bar: ", barCount))
		w.barLabel.Refresh()
		select {
		case <-changeBpmChan:
			newBeat := newBeat(w, p.bpm)
			select {
			case <-click:
				p.beat = newBeat
				p.tic = newBeat - 100
			}

		default:
			for i, char := range *p.pattern {
				select {
				case <-click:
					go func() {
						if w.clickCheck.Checked {
							switch {
							case i == 0 && w.accentDownbeatCheck.Checked:
								playClick(buf.clickDownBeat)
							case w.doubletimeCheck.Checked && i%4 == 0:
								playClick(buf.click)
							case !w.doubletimeCheck.Checked && i%2 == 0:
								playClick(buf.click)
							}
						}
					}()
					go playPattern(char, w, buf.on, buf.filler, buf.off)
				case <-stopPlayChan:
					return
				}
			}
		}
	}
}

func newBeat(w *Widgets, bpm int) int {
	if w.doubletimeCheck.Checked {
		bpm *= 4
	} else {
		bpm *= 2
	}
	return 6000000 / bpm
}

func stop(p *Parameters) {
	stopPlayChan <- struct{}{}
	stopClockChan <- struct{}{}
	p.isPlaying = false
}

func clock(p *Parameters, click *chan struct{}) {
	masterClock := time.NewTicker(10000 * time.Nanosecond)
	for {
		select {
		case <-masterClock.C:
			p.tic++
			if p.tic == p.beat {
				*click <- struct{}{}
				p.tic = 0
			}
		case <-stopClockChan:
			masterClock.Stop()
			return
		}
	}
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
