package rhythmgenerator

func updateButtonStatePlay(w *Widgets) {
	w.invertLeftButton.Enable()
	w.invertRightButton.Enable()
	w.playButton.Disable()
	w.stopButton.Enable()
	w.fillCheckbox.Enable()
	w.omitFillsCheckbox.Enable()
}

func updateButtonStateStop(w *Widgets) {
	w.playButton.Enable()
	w.invertLeftButton.Enable()
	w.invertRightButton.Enable()
	w.stopButton.Disable()
	w.barLabel.SetText("Stopped")
	w.algorithmType.Enable()
	w.removeSymmetryCheckbox.Enable()
}

func initialButtonState(w *Widgets) {
	w.invertLeftButton.Disable()
	w.invertRightButton.Disable()
	w.playButton.Enable()
	w.stopButton.Disable()
	w.algorithmType.Enable()
	w.doubletimeCheckbox.Enable()
	w.omitFillsCheckbox.Disable()
	w.fillCheckbox.Disable()
	w.removeSymmetryCheckbox.Enable()
	w.inversionStatusLabel.SetText("Inversion Status: 0")
}

func filledOk(w *Widgets, enable bool) {
	if enable {
		w.fillOk.SetText("Fill Ok!")
	} else {
		w.fillOk.SetText("Not Filled!")
	}
}

func rsOk(w *Widgets, ok bool) {
	if ok {
		w.RsOk.SetText("Rs: Ok")
	} else {
		w.RsOk.SetText("Rs: n/a")
	}
}

func (w *Widgets) update(pattern string) {
	if len(pattern) > 50 {
		pattern = pattern[:47] + "..."
	}
	w.patternLabel.SetText(pattern)
}
