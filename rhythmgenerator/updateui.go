package rhythmgenerator

func updateButtonStatePlay(w *Widgets) {
	w.invertLeftButton.Enable()
	w.invertRightButton.Enable()
	w.playButton.Disable()
	w.stopButton.Enable()
	w.fillCheck.Enable()
	w.omitFillsCheck.Enable()
}

func updateButtonStateStop(w *Widgets) {
	w.playButton.Enable()
	w.invertLeftButton.Enable()
	w.invertRightButton.Enable()
	w.stopButton.Disable()
	w.barLabel.SetText("Stopped")
	w.algorithmTypeCheck.Enable()
	w.removeSymmetryCheck.Enable()
}

func initialButtonState(w *Widgets) {
	w.invertLeftButton.Disable()
	w.invertRightButton.Disable()
	w.playButton.Enable()
	w.stopButton.Disable()
	w.algorithmTypeCheck.Enable()
	w.doubletimeCheck.Enable()
	w.omitFillsCheck.Disable()
	w.fillCheck.Disable()
	w.removeSymmetryCheck.Enable()
	w.inversionLabel.SetText("Inversion Status: 0")
}

func filledOk(w *Widgets, enable bool) {
	if enable {
		w.fillStatus.SetText("Fill Ok!")
	} else {
		w.fillStatus.SetText("Not Filled!")
	}
}

func rsOk(w *Widgets, ok bool) {
	if ok {
		w.RsStatus.SetText("Rs: Ok")
	} else {
		w.RsStatus.SetText("Rs: n/a")
	}
}

func (w *Widgets) updatePatternLabel(pattern string) {
	if len(pattern) > 50 {
		pattern = pattern[:47] + "..."
	}
	w.patternLabel.SetText(pattern)
}
