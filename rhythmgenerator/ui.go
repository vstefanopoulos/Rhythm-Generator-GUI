package rhythmgenerator

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type widgets struct {
	stepsInput             *widget.Entry
	beatsInput             *widget.Entry
	bpmInput               *widget.Entry
	doubletimeCheckbox     *widget.Check
	playOffsetsCheckbox    *widget.Check
	omitFillsCheckbox      *widget.Check
	playButton             *widget.Button
	stopButton             *widget.Button
	invertRightButton      *widget.Button
	invertLeftButton       *widget.Button
	inversionStatusLabel   *widget.Label
	bar                    *widget.Label
	genPattern             *widget.Label
	algCheckbox            *widget.Check
	fillCheckbox           *widget.Check
	clickCheckbox          *widget.Check
	accentDownbeatCheck    *widget.Check
	removeSymmetryCheckbox *widget.Check
	inversionStatus        int
}

type par struct {
	steps     int
	beats     int
	bpm       int
	euclidean string
	custom    string
	pattern   *string
}

type prev struct {
	stepsInput string
	beatsInput string
	bpmInput   string
}

func Ui() {
	w := &widgets{}
	prev := &prev{}
	par := &par{}
	par.pattern = new(string)

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
	w.accentDownbeatCheck = widget.NewCheck("Accent DownBeat", func(value bool) {})
	w.playOffsetsCheckbox = widget.NewCheck("Play Offsets", func(value bool) {})

	w.algCheckbox = widget.NewCheck("Custom Algorithm", func(value bool) {
		if value {
			chooseCustom(w, par)
		} else {
			chooseEuclidean(w, par)
		}
	})
	w.removeSymmetryCheckbox = widget.NewCheck("Remove Symetry", func(value bool) {
		if value {
			removeSymmetry(w, *par.pattern, par)
		} else {
			fallBack(w, par)
		}
	})

	w.fillCheckbox = widget.NewCheck("Fill Steps", func(value bool) {
		if *par.pattern != "" {
			switch value {
			case true:
				fillSteps(w, par.pattern)
			case false:
				undofillSteps(w, par.pattern)
			}
		}
	})
	w.omitFillsCheckbox = widget.NewCheck("Omit Fills", func(value bool) {})

	w.inversionStatusLabel = widget.NewLabel("")
	w.genPattern = widget.NewLabel("")
	w.bar = widget.NewLabel("")

	w.invertRightButton = widget.NewButton("Invert Right", func() {
		invertPattern(par.pattern, w, true)
	})

	w.invertLeftButton = widget.NewButton("Invert Left", func() {
		invertPattern(par.pattern, w, false)
	})

	w.playButton = widget.NewButton("Play", func() {
		if changedInput(w, prev) {
			var e *Error
			e = callGenerators(w, par)
			if e != nil {
				e.handleInputErrors(w)
				return
			}
			updatePrev(w, prev)
		}
		updateButtonStatePlay(w)
		w.genPattern.SetText(*par.pattern)
		go play(par, w)
	})

	w.stopButton = widget.NewButton("Stop", func() {
		stop()
		updateButtonStateStop(w)
	})
	initialButtonState(w)

	inputBoxCol := container.NewVBox(w.stepsInput, w.beatsInput, w.bpmInput)
	tempoBoxesRow := container.NewHBox(w.doubletimeCheckbox, w.clickCheckbox, w.accentDownbeatCheck, w.playOffsetsCheckbox)
	playStopCol := container.NewVBox(w.playButton, w.stopButton)
	algBoxesRow := container.NewHBox(w.algCheckbox, w.removeSymmetryCheckbox, w.fillCheckbox, w.omitFillsCheckbox)
	invertButtonRow := container.NewHBox(w.invertLeftButton, w.invertRightButton)
	PatBarRow := container.NewHBox(w.genPattern, w.bar)
	allBoxes := container.NewVBox(banner, inputBoxCol, tempoBoxesRow, playStopCol, algBoxesRow,
		invertButtonRow, w.inversionStatusLabel, PatBarRow)
	content := container.NewHBox(allBoxes)
	window.SetContent(content)
	window.ShowAndRun()
}
