package rhythmgenerator

func handleErrors(w *Widgets, p *Parameters, prev *PreviousState) *Error {
	var e *Error
	e = convertInput(w, p)
	updatePrev(w, prev)
	if e != nil {
		initialButtonState(w)
		w.inversionLabel.SetText(e.Message)
		if e.Solution != "" {
			w.patternLabel.SetText(e.Solution)
		}
		if p.isPlaying {
			stop(p)
		}
		return e
	}
	w.updateInversionLabel(p.inversionDegree)
	w.updatePatternLabel(*p.pattern)
	return nil
}
