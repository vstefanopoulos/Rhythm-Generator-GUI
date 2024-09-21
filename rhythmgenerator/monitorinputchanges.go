package rhythmgenerator

import "fmt"

func changedInput(w *widgets, prev *prev) bool {
	if prev.stepsInput != w.stepsInput.Text || prev.beatsInput != w.beatsInput.Text ||
		prev.bpmInput != w.bpmInput.Text || prev.algCheckBox != w.algCheckbox.Checked ||
		prev.fillCheckbox != w.fillCheckbox.Checked || prev.removeSymetryCheckbox != w.removeSymetryCheckbox.Checked {
		fmt.Println("changed input")
		return true
	}
	return false
}

func updatePrev(w *widgets, prev *prev) {
	prev.stepsInput = w.stepsInput.Text
	prev.beatsInput = w.beatsInput.Text
	prev.bpmInput = w.bpmInput.Text
	prev.algCheckBox = w.algCheckbox.Checked
	prev.fillCheckbox = w.fillCheckbox.Checked
	prev.removeSymetryCheckbox = w.removeSymetryCheckbox.Checked
	fmt.Println("Prev updated")
}
