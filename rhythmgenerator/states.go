package rhythmgenerator

import "fyne.io/fyne/v2/widget"

func updateButtonStates(isPlaying bool, playButton, stopButton, invertLeftButton, invertRightButton *widget.Button, bar *widget.Label) {
	if isPlaying {
		invertLeftButton.Disable()
		invertRightButton.Disable()
		playButton.Disable()
		stopButton.Enable()
	} else {
		playButton.Enable()
		invertLeftButton.Enable()
		invertRightButton.Enable()
		stopButton.Disable()
		bar.SetText("Stopped")
	}
}
