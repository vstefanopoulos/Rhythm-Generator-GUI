package rhythmgenerator

func prepForPlay(w *Widgets, p *Parameters, prev *PreviousState) {
	if changedInput(w, prev) {
		var e *Error
		e = callGenerators(w, p)
		if e != nil {
			e.handleInputErrors(w)
			return
		}
		updatePrev(w, prev)
	}
	isPlaying = true
	updateButtonStatePlay(w)
	w.updatePatternLabel(*p.pattern)
}
