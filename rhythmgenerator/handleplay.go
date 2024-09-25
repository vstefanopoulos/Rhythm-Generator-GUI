package rhythmgenerator

func handlePlay(w *Widgets, p *Parameters, prev *PreviousState) {
	if changedInput(w, prev) {
		var e *Error
		e = callGenerators(w, p)
		if e != nil {
			e.handleInputErrors(w)
			return
		}
		updatePrev(w, prev)
	}
	updateButtonStatePlay(w)
	w.genPattern.SetText(*p.pattern)
}
