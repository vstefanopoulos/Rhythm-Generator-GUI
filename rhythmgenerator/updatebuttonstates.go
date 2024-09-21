package rhythmgenerator

import "fmt"

func updateButtonStates(isPlaying bool, w *widgets) {
	if isPlaying {
		w.invertLeftButton.Disable()
		w.invertRightButton.Disable()
		w.playButton.Disable()
		w.stopButton.Enable()
		w.inversionStatusLabel.SetText(fmt.Sprintf("Inversion Status: %v", w.inversionStatus))
		w.algCheckbox.Disable()
		w.doubletimeCheckbox.Disable()
		w.fillCheckbox.Disable()
		w.removeSymetryCheckbox.Disable()
	} else {
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
}
