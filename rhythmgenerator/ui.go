package rhythmgenerator

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var isPlaying bool

type Widgets struct {
	stepsInput          *widget.Entry
	beatsInput          *widget.Entry
	bpmInput            *widget.Entry
	playButton          *widget.Button
	stopButton          *widget.Button
	invertRightButton   *widget.Button
	invertLeftButton    *widget.Button
	inversionLabel      *widget.Label
	barLabel            *widget.Label
	patternLabel        *widget.Label
	fillStatus          *widget.Label
	RsStatus            *widget.Label
	doubletimeCheck     *widget.Check
	omitOffsetsCheck    *widget.Check
	omitFillsCheck      *widget.Check
	algorithmTypeCheck  *widget.Check
	fillCheck           *widget.Check
	clickCheck          *widget.Check
	accentDownbeatCheck *widget.Check
	removeSymmetryCheck *widget.Check
}

type Parameters struct {
	steps           int
	beats           int
	bpm             int
	euclidean       string
	custom          string
	pattern         *string
	inversionDegree int
}

type PreviousState struct {
	stepsInput string
	beatsInput string
	bpmInput   string
}

func Ui() {
	w := &Widgets{}
	prev := &PreviousState{}
	p := &Parameters{}
	p.pattern = new(string)

	RGgui := app.New()
	window := RGgui.NewWindow("Rhythm Generator")

	banner := canvas.NewImageFromFile("./banner.png")
	banner.FillMode = canvas.ImageFillOriginal

	w.stepsInput = widget.NewEntry()
	w.stepsInput.SetPlaceHolder("Steps")
	w.stepsInput.OnSubmitted = func(content string) {
		if isPlaying {
			prepForPlay(w, p, prev)
		}

	}

	w.beatsInput = widget.NewEntry()
	w.beatsInput.SetPlaceHolder("Beats")
	w.beatsInput.OnSubmitted = func(content string) {
		if isPlaying {
			prepForPlay(w, p, prev)
		}
	}

	w.bpmInput = widget.NewEntry()
	w.bpmInput.SetPlaceHolder("BPM")
	w.bpmInput.OnSubmitted = func(content string) {
		if isPlaying {
			var err *Error
			_, _, p.bpm, err = convertInput(w)
			if err != nil {
				stop()
				err.handleInputErrors(w)
			} else {
				changeBpmChan <- struct{}{}
			}
		}
	}

	w.doubletimeCheck = widget.NewCheck("Double Time", func(value bool) {
		if isPlaying {
			changeBpmChan <- struct{}{}
		}
	})

	w.clickCheck = widget.NewCheck("Click", func(value bool) {})
	w.accentDownbeatCheck = widget.NewCheck("Accent DownBeat", func(value bool) {})
	w.omitOffsetsCheck = widget.NewCheck("Omit Offsets", func(value bool) {})

	w.algorithmTypeCheck = widget.NewCheck("Custom Algorithm", func(value bool) {
		if value {
			chooseAlgorithm(w, p, true)
		} else {
			chooseAlgorithm(w, p, false)
		}
	})

	w.removeSymmetryCheck = widget.NewCheck("Remove Symetry", func(value bool) {
		if value {
			removeSymmetry(w, *p.pattern, p)
		} else {
			fallBack(w, p)
		}
	})

	w.fillCheck = widget.NewCheck("Fill Steps", func(value bool) {
		if *p.pattern != "" {
			switch value {
			case true:
				fillSteps(w, p, p.pattern)
			case false:
				undofillSteps(w, p.pattern)
			}
		}
	})

	w.omitFillsCheck = widget.NewCheck("Omit Fills", func(value bool) {})

	w.inversionLabel = widget.NewLabel("")
	w.fillStatus = widget.NewLabel("Fill")
	w.RsStatus = widget.NewLabel("Rs")
	w.patternLabel = widget.NewLabel("")
	w.barLabel = widget.NewLabel("")

	w.invertRightButton = widget.NewButton("Invert Right", func() {
		invertPattern(p.pattern, w, p, true)
	})

	w.invertLeftButton = widget.NewButton("Invert Left", func() {
		invertPattern(p.pattern, w, p, false)
	})

	w.playButton = widget.NewButton("Play", func() {
		prepForPlay(w, p, prev)
		go play(p, w)
	})

	w.stopButton = widget.NewButton("Stop", func() {
		stop()
		updateButtonStateStop(w)
	})

	initialButtonState(w)

	inputBoxCol := container.NewVBox(w.stepsInput, w.beatsInput, w.bpmInput)
	tempoBoxesRow := container.NewHBox(w.doubletimeCheck, w.clickCheck, w.accentDownbeatCheck, w.omitOffsetsCheck)
	playStopCol := container.NewVBox(w.playButton, w.stopButton)
	algBoxesRow := container.NewHBox(w.algorithmTypeCheck, w.removeSymmetryCheck, w.fillCheck, w.omitFillsCheck)
	invertButtonRow := container.NewHBox(w.invertLeftButton, w.invertRightButton, w.inversionLabel, w.RsStatus, w.fillStatus)
	PatBarRow := container.NewHBox(w.patternLabel, w.barLabel)
	allBoxes := container.NewVBox(banner, inputBoxCol, tempoBoxesRow, playStopCol, algBoxesRow,
		invertButtonRow, PatBarRow)
	content := container.NewHBox(allBoxes)
	window.SetContent(content)
	window.ShowAndRun()
}
