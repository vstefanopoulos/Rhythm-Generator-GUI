package rhythmgenerator

func changedInput(w *widgets, prev *prev) bool {
	if prev.stepsInput != w.stepsInput.Text || prev.beatsInput != w.beatsInput.Text ||
		prev.bpmInput != w.bpmInput.Text {
		return true
	}
	return false
}

func updatePrev(w *widgets, prev *prev) {
	prev.stepsInput = w.stepsInput.Text
	prev.beatsInput = w.beatsInput.Text
	prev.bpmInput = w.bpmInput.Text
}
