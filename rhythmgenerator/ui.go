package rhythmgenerator

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type widgets struct {
	stepsInput            *widget.Entry
	beatsInput            *widget.Entry
	bpmInput              *widget.Entry
	playButton            *widget.Button
	stopButton            *widget.Button
	invertRightButton     *widget.Button
	invertLeftButton      *widget.Button
	inversionStatusLabel  *widget.Label
	bar                   *widget.Label
	genPattern            *widget.Label
	algCheckbox           *widget.Check
	fillCheckbox          *widget.Check
	removeSymetryCheckbox *widget.Check
	inversionStatus       int
}

type prev struct {
	stepsInput            string
	beatsInput            string
	bpmInput              string
	algCheckBox           bool
	fillCheckbox          bool
	removeSymetryCheckbox bool
}

func Ui() {
	var pattern string
	var bpm int
	w := &widgets{}
	prev := &prev{}

	RGgui := app.New()
	window := RGgui.NewWindow("Rhythm Generator")
	window.Resize(fyne.NewSize(400, 300))

	w.stepsInput = widget.NewEntry()
	w.stepsInput.SetPlaceHolder("Steps")

	w.beatsInput = widget.NewEntry()
	w.beatsInput.SetPlaceHolder("Beats")

	w.bpmInput = widget.NewEntry()
	w.bpmInput.SetPlaceHolder("BPM")

	w.algCheckbox = widget.NewCheck("Custom Algorithm", func(value bool) {})

	w.fillCheckbox = widget.NewCheck("Fill Steps", func(value bool) {})

	w.removeSymetryCheckbox = widget.NewCheck("Remove Symetry", func(value bool) {})

	w.inversionStatusLabel = widget.NewLabel("")

	w.bar = widget.NewLabel("")

	w.genPattern = widget.NewLabel("")

	w.invertRightButton = widget.NewButton("Invert Right", func() {
		pattern = invertPattern(pattern, w, true)
	})

	w.invertLeftButton = widget.NewButton("Invert Left", func() {
		pattern = invertPattern(pattern, w, false)
	})

	w.playButton = widget.NewButton("Play", func() {
		updateButtonStates(true, w)
		if pattern == "" || changedInput(w, prev) {
			var e *Error
			pattern, bpm, e = callGenerators(w, prev)
			if e != nil {
				e.handleInputErrors(w)
				return
			}
			updatePrev(w, prev)
		}
		w.genPattern.SetText(pattern)
		go play(pattern, bpm, w)
	})

	w.stopButton = widget.NewButton("Stop", func() {
		stop()
		updateButtonStates(false, w)

	})
	initialButtonState(w)

	inputBoxCol := container.NewVBox(w.stepsInput, w.beatsInput, w.bpmInput)
	playStopCol := container.NewVBox(w.playButton, w.stopButton)
	invertButtonRow := container.NewHBox(w.invertLeftButton, w.invertRightButton)
	checkBoxesRow := container.NewHBox(w.algCheckbox, w.fillCheckbox, w.removeSymetryCheckbox)
	PatBarRow := container.NewHBox(w.genPattern, w.bar)

	content := container.NewVBox(inputBoxCol, playStopCol, invertButtonRow, checkBoxesRow, w.inversionStatusLabel, PatBarRow)
	window.SetContent(content)
	window.ShowAndRun()
}
