package rhythmgenerator

const onSet = 'X'
const offSet = 'o'
const fill = 'x'

func callGenerators(w *widgets) (*string, int, *Error) {
	var pattern string
	steps, beats, bpm, err := convertInput(w)

	if err != nil {
		return nil, 0, err
	}
	w.inversionStatus = 0

	if w.algCheckbox.Checked {
		pattern = customGenerate(steps, beats)
	} else {
		pattern = euclideanGenerate(steps, beats)
	}

	if w.removeSymetryCheckbox.Checked {
		newPattern, generatedAsymetrical := removeSymetry(pattern, steps, beats)
		if generatedAsymetrical {
			pattern = newPattern
		} else {
			w.removeSymetryCheckbox.SetChecked(false)
		}
	}

	if steps/beats > 1 && w.fillCheckbox.Checked {
		newPattern, filledSteps := fillSteps(pattern)
		if filledSteps {
			pattern = newPattern
			filledButtonState(w, true)
		} else {
			filledButtonState(w, false)
		}
	}
	return &pattern, bpm, nil
}
