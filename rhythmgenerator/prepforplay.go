package rhythmgenerator

func prepForPlay(w *Widgets, p *Parameters, prev *PreviousState) {
	if changedInput(w, prev) {
		var e *Error
		e = callGenerators(w, p)
		if e != nil {
			e.handleInputErrors(w, p)
			return
		}
		updatePrev(w, prev)
	}
	updateButtonStatePlay(w)
	w.updatePatternLabel(*p.pattern)
}
