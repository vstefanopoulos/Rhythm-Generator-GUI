package rhythmgenerator

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gopxl/beep"
)

type Widgets struct {
	stepsInput          *widget.Entry
	beatsInput          *widget.Entry
	bpmInput            *widget.Entry
	playButton          *widget.Button
	stopButton          *widget.Button
	invertRightButton   *widget.Button
	invertLeftButton    *widget.Button
	inversionLabel      *widget.Label
	barLabel            *canvas.Text
	patternLabel        *canvas.Text
	fillStatus          *canvas.Text
	rsStatus            *canvas.Text
	errLabel            *canvas.Text
	errSolutionLabel    *canvas.Text
	doubletimeCheck     *widget.Check
	muteOffsetsCheck    *widget.Check
	muteFillsCheck      *widget.Check
	algorithmTypeCheck  *widget.Check
	fillCheck           *widget.Check
	clickCheck          *widget.Check
	accentDownbeatCheck *widget.Check
	removeSymmetryCheck *widget.Check
	addKick             *widget.Check
}

type Parameters struct {
	steps           int
	beats           int
	bpm             int
	inversionDegree int
	beat            int
	clockBuffer     int
	euclidean       string
	custom          string
	pattern         *string
	isPlaying       bool
}

type Buffer struct {
	snr           *beep.Buffer
	side          *beep.Buffer
	hh            *beep.Buffer
	clickDownBeat *beep.Buffer
	click         *beep.Buffer
	kick          *beep.Buffer
}

func Ui() {
	w := &Widgets{}
	p := &Parameters{}
	p.pattern = new(string)
	buf := &Buffer{
		kick:          makeBuffer("./wav/kick.wav"),
		snr:           makeBuffer("./wav/rim.wav"),
		side:          makeBuffer("./wav/side.wav"),
		hh:            makeBuffer("./wav/hh.wav"),
		clickDownBeat: makeBuffer("./wav/ride_bell.wav"),
		click:         makeBuffer("./wav/tambourine.wav"),
	}
	RGgui := app.New()
	window := RGgui.NewWindow("Rhythm Generator")

	banner := canvas.NewImageFromFile("./banner.png")
	banner.FillMode = canvas.ImageFillOriginal

	w.stepsInput = widget.NewEntry()
	w.stepsInput.SetPlaceHolder("Steps")
	w.stepsInput.OnSubmitted = func(content string) {
		e := handleErrors(w, p)
		if e != nil {
			return
		}
		callGenerators(w, p)
	}

	w.beatsInput = widget.NewEntry()
	w.beatsInput.SetPlaceHolder("Beats")
	w.beatsInput.OnSubmitted = func(content string) {
		e := handleErrors(w, p)
		if e != nil {
			return
		}
		callGenerators(w, p)
	}

	w.bpmInput = widget.NewEntry()
	w.bpmInput.SetPlaceHolder("BPM")
	w.bpmInput.OnSubmitted = func(content string) {
		e := handleErrors(w, p)
		if e != nil {
			return
		}
		if p.isPlaying {
			changeBpmChan <- struct{}{}
		} else {
			callGenerators(w, p)
		}
	}

	// C H E C K  B O X E S - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	w.doubletimeCheck = widget.NewCheck("Double Time", func(value bool) {
		e := handleErrors(w, p)
		if e != nil {
			return
		}
		if p.isPlaying {
			go func() {
				changeBpmChan <- struct{}{}
			}()
		}
	})

	w.algorithmTypeCheck = widget.NewCheck("Custom Algorithm", func(value bool) {
		e := handleErrors(w, p)
		if e != nil {
			return
		}
		callGenerators(w, p)
	})

	w.removeSymmetryCheck = widget.NewCheck("Remove Symetry", func(value bool) {
		e := handleErrors(w, p)
		if e != nil {
			return
		}
		callGenerators(w, p)
	})

	w.fillCheck = widget.NewCheck("Fill Steps", func(value bool) {
		e := handleErrors(w, p)
		if e != nil {
			return
		}
		callGenerators(w, p)
	})

	w.addKick = widget.NewCheck("Add Kick", func(value bool) {
		e := handleErrors(w, p)
		if e != nil {
			return
		}
		callGenerators(w, p)
	})

	w.clickCheck = widget.NewCheck("Click", func(value bool) {})
	w.accentDownbeatCheck = widget.NewCheck("Accent DownBeat", func(value bool) {})
	w.muteOffsetsCheck = widget.NewCheck("Mute Offsets", func(value bool) {})
	w.muteFillsCheck = widget.NewCheck("Mute Fills", func(value bool) {})

	// B U T T O N S - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	w.invertRightButton = widget.NewButton("Invert Right", func() {
		invertRight(p.pattern, w, p)
	})

	w.invertLeftButton = widget.NewButton("Invert Left", func() {
		invertLeft(p.pattern, w, p)
	})

	w.playButton = widget.NewButton("Play", func() {
		e := handleErrors(w, p)
		if e != nil {
			return
		}
		callGenerators(w, p)
		updateButtonStatePlay(w)
		go play(p, w, buf)
	})

	w.stopButton = widget.NewButton("Stop", func() {
		stop(p)
		updateButtonStateStop(w)
	})

	// S T A T U S  L A B E L S - - - - - - - - - - - - - - - - - - - - - - - -
	w.inversionLabel = widget.NewLabel("")

	w.fillStatus = canvas.NewText("", color.RGBA{0, 255, 0, 200})
	w.rsStatus = canvas.NewText("", color.RGBA{0, 255, 0, 200})

	w.barLabel = canvas.NewText("", color.RGBA{255, 0, 0, 200})
	w.barLabel.TextSize = 15
	w.patternLabel = canvas.NewText("", color.RGBA{0, 255, 0, 200})
	w.patternLabel.TextSize = 15

	// E R R  L A B E L S - - - - - - - - - - - - - - - - - - - - -
	w.errLabel = canvas.NewText("", color.RGBA{255, 0, 0, 255})
	w.errSolutionLabel = canvas.NewText("", color.RGBA{255, 0, 0, 255})

	initialButtonState(w)

	// C O N T A I N E R S - - - - - - - - - - - - - - - - - - - - - - - - -
	inputBoxCol := container.NewVBox(w.beatsInput, w.stepsInput, w.bpmInput)
	tempoBoxesRow := container.NewHBox(w.doubletimeCheck, w.clickCheck, w.accentDownbeatCheck, w.muteOffsetsCheck)
	playStopCol := container.NewVBox(w.playButton, w.stopButton)
	algMuteRow := container.NewHBox(w.algorithmTypeCheck, w.removeSymmetryCheck, w.rsStatus)
	simFillRow := container.NewHBox(w.fillCheck, w.fillStatus, w.muteFillsCheck, w.addKick)
	invertButtonRow := container.NewHBox(w.invertLeftButton, w.invertRightButton, w.inversionLabel)
	PatBarRow := container.NewHBox(w.patternLabel, w.barLabel)
	errRow := container.NewVBox(w.errLabel, w.errSolutionLabel)
	allBoxes := container.NewVBox(banner, inputBoxCol, tempoBoxesRow, playStopCol, algMuteRow, simFillRow,
		invertButtonRow, PatBarRow, errRow)
	content := container.NewHBox(allBoxes)

	window.SetContent(content)
	window.ShowAndRun()
}
