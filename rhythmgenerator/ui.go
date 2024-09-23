package rhythmgenerator

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type widgets struct {
	stepsInput            *widget.Entry
	beatsInput            *widget.Entry
	bpmInput              *widget.Entry
	doubletimeCheckbox    *widget.Check
	playOffsetsCheckbox   *widget.Check
	playFillsCheckbox     *widget.Check
	playButton            *widget.Button
	stopButton            *widget.Button
	invertRightButton     *widget.Button
	invertLeftButton      *widget.Button
	inversionStatusLabel  *widget.Label
	bar                   *widget.Label
	genPattern            *widget.Label
	algCheckbox           *widget.Check
	fillCheckbox          *widget.Check
	clickCheckbox         *widget.Check
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
	doubletimeCheckbox    bool
}

func Ui() {
	var pattern *string = new(string)
	var bpm int
	w := &widgets{}
	prev := &prev{}

	RGgui := app.New()
	window := RGgui.NewWindow("Rhythm Generator")

	banner := canvas.NewImageFromFile("./banner.png")
	banner.FillMode = canvas.ImageFillOriginal

	w.stepsInput = widget.NewEntry()
	w.stepsInput.SetPlaceHolder("Steps")

	w.beatsInput = widget.NewEntry()
	w.beatsInput.SetPlaceHolder("Beats")

	w.bpmInput = widget.NewEntry()
	w.bpmInput.SetPlaceHolder("BPM")

	w.doubletimeCheckbox = widget.NewCheck("Double Time", func(value bool) {})
	w.clickCheckbox = widget.NewCheck("Click", func(value bool) {})
	w.playOffsetsCheckbox = widget.NewCheck("Play Offsets", func(value bool) {})

	w.algCheckbox = widget.NewCheck("Custom Algorithm", func(value bool) {})
	w.removeSymetryCheckbox = widget.NewCheck("Remove Symetry", func(value bool) {})

	w.fillCheckbox = widget.NewCheck("Fill Steps", func(value bool) {
		if *pattern != "" {
			switch value {
			case true:
				go fillSteps(w, pattern)
			case false:
				go undofillSteps(w, pattern)
			}
		}
	})
	w.playFillsCheckbox = widget.NewCheck("Play Fills", func(value bool) {})

	w.inversionStatusLabel = widget.NewLabel("")
	w.genPattern = widget.NewLabel("")
	w.bar = widget.NewLabel("")

	w.invertRightButton = widget.NewButton("Invert Right", func() {
		go invertPattern(pattern, w, true)
	})

	w.invertLeftButton = widget.NewButton("Invert Left", func() {
		go invertPattern(pattern, w, false)
	})

	w.playButton = widget.NewButton("Play", func() {
		if changedInput(w, prev) {
			var e *Error
			pattern, bpm, e = callGenerators(w)
			if e != nil {
				e.handleInputErrors(w)
				return
			}
			updatePrev(w, prev)
		}
		updateButtonStatePlay(w)
		w.genPattern.SetText(*pattern)
		go play(pattern, bpm, w)
	})

	w.stopButton = widget.NewButton("Stop", func() {
		stop()
		updateButtonStateStop(w)

	})
	initialButtonState(w)

	inputBoxCol := container.NewVBox(w.stepsInput, w.beatsInput, w.bpmInput)
	tempoBoxesRow := container.NewHBox(w.doubletimeCheckbox, w.clickCheckbox, w.playOffsetsCheckbox)
	playStopCol := container.NewVBox(w.playButton, w.stopButton)
	algBoxesRow := container.NewHBox(w.algCheckbox, w.removeSymetryCheckbox)
	fillBoxesRow := container.NewHBox(w.fillCheckbox, w.playFillsCheckbox)
	invertButtonRow := container.NewHBox(w.invertLeftButton, w.invertRightButton)
	PatBarRow := container.NewHBox(w.genPattern, w.bar)
	allBoxes := container.NewVBox(banner, inputBoxCol, tempoBoxesRow, playStopCol, algBoxesRow, fillBoxesRow,
		invertButtonRow, w.inversionStatusLabel, PatBarRow)
	content := container.NewHBox(allBoxes)
	window.SetContent(content)
	window.ShowAndRun()
}
