package rhythmgenerator

import "fyne.io/fyne/v2/widget"

func handleInputErrors(bar, patternInfo *widget.Label, playButton, stopButton *widget.Button) bool {
	if InputError != "" {
		patternInfo.SetText(InputError)
		if InputErrorSolution != "" {
			bar.SetText(InputErrorSolution)
		}
		playButton.Enable()
		stopButton.Disable()
		return true
	} else {
		return false
	}
}
