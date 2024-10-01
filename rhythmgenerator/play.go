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
	p.clockBuffer = p.beat - 10
	click := make(chan struct{})
	go clock(p, &click)
	var barCount int
	var doubleTime bool
	p.isPlaying = true
	for {
		go func() {
			doubleTime = w.doubletimeCheck.Checked
			barCount++
			w.barLabel.Text = (fmt.Sprint("Bar: ", barCount))
			w.barLabel.Refresh()
		}()
		select {
		case <-changeBpmChan:
			newBeat := newBeat(w, p.bpm)
			select {
			case <-click:
				p.beat = newBeat
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
							case doubleTime && i%4 == 0:
								playClick(buf.click)
							case !doubleTime && i%2 == 0:
								playClick(buf.click)
							}
						}
					}()
					go playSpeaker(char, w, buf)
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
	masterClock := time.NewTicker(10 * time.Microsecond)
	for {
		select {
		case <-masterClock.C:
			p.clockBuffer++
			if p.clockBuffer == p.beat {
				*click <- struct{}{}
				p.clockBuffer = 0
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

func playSpeaker(char rune, w *Widgets, buf *Buffer) {
	switch {
	case char == 'X':
		speaker.Play(buf.snr.Streamer(0, buf.snr.Len()))
	case char == 'x':
		if !w.muteFillsCheck.Checked {
			speaker.Play(buf.side.Streamer(0, buf.side.Len()))
		} else {
			return
		}
	case char == 'o':
		if !w.muteOffsetsCheck.Checked {
			speaker.Play(buf.hh.Streamer(0, buf.hh.Len()))
		} else {
			return
		}
	}
}

func playClick(click *beep.Buffer) {
	streamer := click.Streamer(0, click.Len())
	speaker.Play(streamer)
}
