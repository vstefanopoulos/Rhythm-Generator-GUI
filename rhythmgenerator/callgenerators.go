package rhythmgenerator

const onSet = 'X'
const offSet = 'o'
const fill = 'x'

func callGenerators(w *widgets) (string, int, *Error) {
	var pattern string
	steps, beats, bpm, err := convertInput(w)

	if err != nil {
		return "", 0, err
	}
	w.inversionStatus = 0

	if w.algCheckbox.Checked {
		pattern = customGenerate(steps, beats)
	} else {
		pattern = euclideanGenerate(steps, beats)
	}

	if w.removeSymetryCheckbox.Checked && steps > 9 && beats > 3 && steps%beats != 0 {
		newPattern, isSymetrical := removeSymetry(pattern, steps)
		if isSymetrical {
			pattern = newPattern
		}
	}

	if steps/beats > 1 && w.fillCheckbox.Checked {
		newPattern, filledSteps := fillSteps(pattern)
		if filledSteps {
			pattern = newPattern
			w.playFillsCheckbox.Enable()
		} else {
			w.playFillsCheckbox.SetChecked(false)
			w.playFillsCheckbox.Disable()
		}
	}
	return pattern, bpm, nil
}
