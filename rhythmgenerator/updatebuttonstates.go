package rhythmgenerator

import "fmt"

func updateButtonStatePlay(w *widgets) {

	w.invertLeftButton.Enable()
	w.invertRightButton.Enable()
	w.playButton.Disable()
	w.stopButton.Enable()
	w.inversionStatusLabel.SetText(fmt.Sprintf("Inversion Status: %v", w.inversionStatus))
	w.doubletimeCheckbox.Disable()
	w.fillCheckbox.Enable()
	w.omitFillsCheckbox.Enable()
}

func updateButtonStateStop(w *widgets) {
	w.playButton.Enable()
	w.invertLeftButton.Enable()
	w.invertRightButton.Enable()
	w.stopButton.Disable()
	w.bar.SetText("Stopped")
	w.algCheckbox.Enable()
	w.doubletimeCheckbox.Enable()
	w.fillCheckbox.Enable()
	w.omitFillsCheckbox.Enable()
	w.removeSymmetryCheckbox.Enable()
}

func initialButtonState(w *widgets) {
	w.invertLeftButton.Disable()
	w.invertRightButton.Disable()
	w.playButton.Enable()
	w.stopButton.Disable()
	w.algCheckbox.Enable()
	w.doubletimeCheckbox.Enable()
	w.omitFillsCheckbox.Disable()
	w.fillCheckbox.Disable()
	w.removeSymmetryCheckbox.Enable()
	w.playOffsetsCheckbox.SetChecked(true)
}

func filledButtonState(w *widgets, enable bool) {
	if enable {
		w.omitFillsCheckbox.Enable()
	} else {
		w.omitFillsCheckbox.SetChecked(false)
		w.omitFillsCheckbox.Disable()
		w.fillCheckbox.SetChecked(false)
		w.fillCheckbox.Disable()
	}
}
