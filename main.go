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
	var pattern string
	var T string
	var bpm int

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

	bpmInput := widget.NewEntry()
	bpmInput.SetPlaceHolder("BPM")

	playing := widget.NewLabel("")

	bar := widget.NewLabel("")

	playingPattern := widget.NewLabel("")

	stopChannel := make(chan struct{})
	stopChannel2 := make(chan struct{})
	var playButton *widget.Button
	var stopButton *widget.Button
	var invertButton *widget.Button

	stopButton = widget.NewButton("Stop", func() {
		go func() {
			stopChannel <- struct{}{}
		}()
		go func() {
			stopChannel2 <- struct{}{}
		}()
		invertButton.Enable()
		playButton.Enable()
		stopButton.Disable()
		bar.SetText("Stopped")
	})

	invertButton = widget.NewButton("Invert Pattern", func() {
		if pattern != "" {
			pattern = pattern[1:] + pattern[0:1]
			playingPattern.SetText(pattern)
		}
	})
	invertButton.Disable()

	playButton = widget.NewButton("Play", func() {
		invertButton.Disable()
		playButton.Disable()
		stopButton.Enable()
		args[0] = steps.Text
		args[1] = beats.Text
		args[2] = bpmInput.Text
		args[3] = algType
		args[4] = fill
		if pattern == "" {
			pattern, T, bpm = rhythmgenerator.RGMain(args)
		}
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

	content := container.NewVBox(steps, beats, bpmInput, playButton, stopButton, invertButton, algCheckbox, fillCheckbox, playing, bar, playingPattern)
	window.SetContent(content)
	window.ShowAndRun()
}
