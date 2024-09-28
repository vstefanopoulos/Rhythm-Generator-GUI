package rhythmgenerator

import "fmt"

func updateButtonStatePlay(w *Widgets) {
	w.playButton.Disable()
	w.stopButton.Enable()
	w.muteFillsCheck.Enable()
}

func updateButtonStateStop(w *Widgets) {
	w.playButton.Enable()
	w.stopButton.Disable()
	w.barLabel.Text = "Stopped"
	w.barLabel.Refresh()
}

func initialButtonState(w *Widgets) {
	w.playButton.Enable()
	w.stopButton.Disable()
	w.inversionLabel.SetText("Inversion Status: n/a")
	rsOk(w, false)
	filledOk(w, false)
}

func filledOk(w *Widgets, enable bool) {
	if enable {
		w.fillStatus.Text = "\u2713Ok"
		w.fillStatus.Refresh()
	} else {
		w.fillStatus.Text = ""
		w.fillStatus.Refresh()
	}
}

func rsOk(w *Widgets, ok bool) {
	if ok {
		w.rsStatus.Text = "\u2713Ok "
		w.rsStatus.Refresh()
	} else {
		w.rsStatus.Text = ""
		w.rsStatus.Refresh()
	}
}

func (w *Widgets) updatePatternLabel(pattern string) {
	if len(pattern) > 50 {
		pattern = pattern[:47] + "..."
	}
	w.patternLabel.Text = pattern + "\t"
	w.patternLabel.Refresh()
}

func (w *Widgets) updateInversionLabel(degree int) {
	w.inversionLabel.SetText(fmt.Sprintf("Inversion Status: %v", degree))
}
