package rhythmgenerator

import "fmt"

func updateButtonStatePlay(w *widgets) {

	w.invertLeftButton.Enable()
	w.invertRightButton.Enable()
	w.playButton.Disable()
	w.stopButton.Enable()
	w.inversionStatusLabel.SetText(fmt.Sprintf("Inversion Status: %v", w.inversionStatus))
	w.algCheckbox.Disable()
	w.doubletimeCheckbox.Disable()
	w.removeSymetryCheckbox.Disable()
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
	w.removeSymetryCheckbox.Enable()
}

func initialButtonState(w *widgets) {
	w.invertLeftButton.Disable()
	w.invertRightButton.Disable()
	w.playButton.Enable()
	w.stopButton.Disable()
	w.algCheckbox.Enable()
	w.doubletimeCheckbox.Enable()
	w.playFillsCheckbox.Disable()
	w.removeSymetryCheckbox.Enable()
	w.playOffsetsCheckbox.SetChecked(true)
}

func filledButtonState(w *widgets, enable bool) {
	if enable {
		w.playFillsCheckbox.Enable()
	} else {
		w.playFillsCheckbox.SetChecked(false)
		w.playFillsCheckbox.Disable()
		w.fillCheckbox.Disable()
	}
}
