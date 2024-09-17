package main

import (
	"fmt"
	"rhythmgenerator/rhythmgenerator"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	args := make([]string, 5)

	RGgui := app.New()
	window := RGgui.NewWindow("Rhythm Generator")
	window.Resize(fyne.NewSize(400, 300))

	var algType string = "Euclidean"
	algCheckbox := widget.NewCheck("Use Custom Algorithm", func(value bool) {
		if value {
			algType = "Custom"
		} else {
			algType = "Euclidean"
		}
	})

	var fill string
	fillCheckbox := widget.NewCheck("Fill Steps", func(value bool) {
		if value {
			fill = "fill"
		} else {
			fill = ""
		}
	})

	steps := widget.NewEntry()
	steps.SetPlaceHolder("Steps")

	beats := widget.NewEntry()
	beats.SetPlaceHolder("Beats")

	bpm := widget.NewEntry()
	bpm.SetPlaceHolder("BPM")

	playing := widget.NewLabel("")

	bar := widget.NewLabel("")

	playingPattern := widget.NewLabel("")

	stopChannel := make(chan struct{})
	stopChannel2 := make(chan struct{})
	var playButton *widget.Button
	var stopButton *widget.Button

	stopButton = widget.NewButton("Stop", func() {
		go func() {
			stopChannel <- struct{}{}
		}()
		go func() {
			stopChannel2 <- struct{}{}
		}()
		playButton.Enable()
		stopButton.Disable()
		bar.SetText("Stopped")
	})

	playButton = widget.NewButton("Play", func() {
		playButton.Disable()
		stopButton.Enable()
		args[0] = steps.Text
		args[1] = beats.Text
		args[2] = bpm.Text
		args[3] = algType
		args[4] = fill

		pattern, T, bpm := rhythmgenerator.RGMain(args)
		if rhythmgenerator.InputError != "" {
			playing.SetText(rhythmgenerator.InputError)
			if rhythmgenerator.InputErrorSolution != "" {
				bar.SetText(rhythmgenerator.InputErrorSolution)
			}
			playButton.Enable()
			stopButton.Disable()
			return
		}
		if pattern != "" {
			go func() {
				var wg sync.WaitGroup
				playing.SetText(fmt.Sprintf("Pattern: %v Algorithm", T))
				playingPattern.SetText(pattern)
				var barCount int

				for {
					select {
					case <-stopChannel:
						return
					default:
						barCount += 1

						wg.Add(2)

						go func() {
							defer wg.Done()
							bar.SetText(fmt.Sprintf("Bar: %v", barCount))
						}()

						go func() {
							defer wg.Done()
							rhythmgenerator.PlayPattern(pattern, bpm, stopChannel2)
						}()

						wg.Wait()
					}
				}
			}()
		}
	})

	content := container.NewVBox(steps, beats, bpm, playButton, stopButton, algCheckbox, fillCheckbox, playing, bar, playingPattern)
	window.SetContent(content)
	window.ShowAndRun()
}
