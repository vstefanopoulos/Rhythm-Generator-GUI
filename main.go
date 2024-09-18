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
	algCheckbox := widget.NewCheck("Custom Algorithm", func(value bool) {
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

	patternInfo := widget.NewLabel("")

	bar := widget.NewLabel("")

	genPattern := widget.NewLabel("")

	stopChannel := make(chan struct{})
	stopChannel2 := make(chan struct{})
	var playButton *widget.Button
	var stopButton *widget.Button
	var invertRightButton *widget.Button
	var invertLeftButton *widget.Button

	stopButton = widget.NewButton("Stop", func() {
		go func() {
			stopChannel <- struct{}{}
		}()
		go func() {
			stopChannel2 <- struct{}{}
		}()
		invertLeftButton.Enable()
		invertRightButton.Enable()
		playButton.Enable()
		stopButton.Disable()
		bar.SetText("Stopped")
	})
	stopButton.Disable()

	var inverted bool
	invertRightButton = widget.NewButton("Invert Right", func() {
		if pattern != "" {
			pattern = pattern[len(pattern)-1:] + pattern[0:len(pattern)-1]
			genPattern.SetText(pattern)
			inverted = true
		}
	})
	invertRightButton.Disable()

	invertLeftButton = widget.NewButton("Invert Left", func() {
		if pattern != "" {
			pattern = pattern[1:] + pattern[0:1]
			genPattern.SetText(pattern)
			inverted = true
		}
	})
	invertLeftButton.Disable()

	playButton = widget.NewButton("Play", func() {
		invertLeftButton.Disable()
		invertRightButton.Disable()
		playButton.Disable()
		stopButton.Enable()
		args[0] = steps.Text
		args[1] = beats.Text
		args[2] = bpmInput.Text
		args[3] = algType
		args[4] = fill
		if pattern == "" || !inverted {
			pattern, T, bpm = rhythmgenerator.RGMain(args)
		}
		inverted = false
		if rhythmgenerator.InputError != "" {
			patternInfo.SetText(rhythmgenerator.InputError)
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
				patternInfo.SetText(fmt.Sprintf("Pattern: %v Algorithm", T))
				genPattern.SetText(pattern)
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

	invertButtonRow := container.NewHBox(invertLeftButton, invertRightButton)
	checkBoxesRow := container.NewHBox(algCheckbox, fillCheckbox)
	infoBarRow := container.NewHBox(patternInfo, bar)

	content := container.NewVBox(steps, beats, bpmInput, playButton, stopButton, invertButtonRow, checkBoxesRow, infoBarRow, genPattern)
	window.SetContent(content)
	window.ShowAndRun()
}
