package rhythmgenerator

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Ui() {
	var pattern string
	var bpm int
	var playButton *widget.Button
	var stopButton *widget.Button
	var invertRightButton *widget.Button
	var invertLeftButton *widget.Button

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

	var fill bool
	fillCheckbox := widget.NewCheck("Fill Steps", func(value bool) {
		if value {
			fill = true
		} else {
			fill = false
		}
	})

	var removerSymetry bool
	removerSymetryCheckbox := widget.NewCheck("Remove Symetry", func(value bool) {
		if value {
			removerSymetry = true
		} else {
			removerSymetry = false
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

	stopButton = widget.NewButton("Stop", func() {
		stop()
		updateButtonStates(false, playButton, stopButton, invertLeftButton, invertRightButton, bar)

	})
	stopButton.Disable()

	var inverted bool
	invertRightButton = widget.NewButton("Invert Right", func() {
		pattern = invertPattern(pattern, &inverted, genPattern, true)
	})
	invertRightButton.Disable()

	invertLeftButton = widget.NewButton("Invert Left", func() {
		pattern = invertPattern(pattern, &inverted, genPattern, false)
	})
	invertLeftButton.Disable()

	playButton = widget.NewButton("Play", func() {
		updateButtonStates(true, playButton, stopButton, invertLeftButton, invertRightButton, bar)

		if pattern == "" || !inverted {
			pattern, bpm = callGenerators(steps.Text, beats.Text, bpmInput.Text, algType, fill, removerSymetry)
		}
		inverted = false

		if handleInputErrors(bar, patternInfo, playButton, stopButton) {
			return
		}
		go play(pattern, bpm, bar, algType, patternInfo, genPattern)

	})

	invertButtonRow := container.NewHBox(invertLeftButton, invertRightButton)
	checkBoxesRow := container.NewHBox(algCheckbox, fillCheckbox, removerSymetryCheckbox)
	infoBarRow := container.NewHBox(patternInfo, bar)

	content := container.NewVBox(steps, beats, bpmInput, playButton, stopButton, invertButtonRow, checkBoxesRow, infoBarRow, genPattern)
	window.SetContent(content)
	window.ShowAndRun()
}
