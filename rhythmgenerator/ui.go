package rhythmgenerator

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Ui() {

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

	var playButton *widget.Button
	var stopButton *widget.Button
	var invertRightButton *widget.Button
	var invertLeftButton *widget.Button

	stopButton = widget.NewButton("Stop", func() {
		Stop()
		updateButtonStates(false, playButton, stopButton, invertLeftButton, invertRightButton, bar)

	})
	stopButton.Disable()

	var inverted bool
	invertRightButton = widget.NewButton("Invert Right", func() {
		invertPattern(pattern, &inverted, genPattern, true)
	})
	invertRightButton.Disable()

	invertLeftButton = widget.NewButton("Invert Left", func() {
		invertPattern(pattern, &inverted, genPattern, false)
	})
	invertLeftButton.Disable()

	playButton = widget.NewButton("Play", func() {
		updateButtonStates(true, playButton, stopButton, invertLeftButton, invertRightButton, bar)

		args[0] = steps.Text
		args[1] = beats.Text
		args[2] = bpmInput.Text
		args[3] = algType
		args[4] = fill

		if pattern == "" || !inverted {
			pattern, T, bpm = callGenerators(args)
		}
		inverted = false

		if handleInputErrors(bar, patternInfo, playButton, stopButton) {
			return
		}
		go Play(pattern, bpm, bar, T, patternInfo, genPattern)

	})

	invertButtonRow := container.NewHBox(invertLeftButton, invertRightButton)
	checkBoxesRow := container.NewHBox(algCheckbox, fillCheckbox)
	infoBarRow := container.NewHBox(patternInfo, bar)

	content := container.NewVBox(steps, beats, bpmInput, playButton, stopButton, invertButtonRow, checkBoxesRow, infoBarRow, genPattern)
	window.SetContent(content)
	window.ShowAndRun()
}
