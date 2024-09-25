package rhythmgenerator

func changedInput(w *Widgets, prev *PreviousState) bool {
	if prev.stepsInput != w.stepsInput.Text || prev.beatsInput != w.beatsInput.Text ||
		prev.bpmInput != w.bpmInput.Text {
		return true
	}
	return false
}

func updatePrev(w *Widgets, prev *PreviousState) {
	prev.stepsInput = w.stepsInput.Text
	prev.beatsInput = w.beatsInput.Text
	prev.bpmInput = w.bpmInput.Text
}
