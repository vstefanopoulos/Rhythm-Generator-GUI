package rhythmgenerator

import "fmt"

func updateButtonStates(isPlaying bool, w *widgets) {
	if isPlaying {
		w.invertLeftButton.Disable()
		w.invertRightButton.Disable()
		w.playButton.Disable()
		w.stopButton.Enable()
		w.inversionStatusLabel.SetText(fmt.Sprintf("Inversion Status: %v", w.inversionStatus))
	} else {
		w.playButton.Enable()
		w.invertLeftButton.Enable()
		w.invertRightButton.Enable()
		w.stopButton.Disable()
		w.bar.SetText("Stopped")
	}

}

func initialButtonState(w *widgets) {
	w.invertLeftButton.Disable()
	w.invertRightButton.Disable()
	w.playButton.Enable()
	w.stopButton.Disable()
}
