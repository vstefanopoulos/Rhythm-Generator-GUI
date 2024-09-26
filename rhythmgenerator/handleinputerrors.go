package rhythmgenerator

type Error struct {
	Message  string
	Solution string
}

func (err *Error) handleInputErrors(w *Widgets, p *Parameters) {
	if err != nil {
		w.inversionLabel.SetText(err.Message)
		if err.Solution != "" {
			w.patternLabel.SetText(err.Solution)
		}
		if p.isPlaying {
			stop(p)
		}
		initialButtonState(w)
	}
}
