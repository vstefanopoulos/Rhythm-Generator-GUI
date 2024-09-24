package rhythmgenerator

const onSet = 'X'
const offSet = 'o'
const fill = 'x'

func callGenerators(w *widgets, par *par) *Error {
	var pattern string
	var err *Error
	par.steps, par.beats, par.bpm, err = convertInput(w)

	if err != nil {
		return err
	}
	w.inversionStatus = 0

	if w.algCheckbox.Checked {
		pattern = customGenerate(par.steps, par.beats)
	} else {
		pattern = euclideanGenerate(par.steps, par.beats)
	}

	if w.removeSymetryCheckbox.Checked {
		newPattern, generatedAsymetrical := removeSymetry(pattern, par.steps, par.beats)
		if generatedAsymetrical {
			pattern = newPattern
		} else {
			w.removeSymetryCheckbox.SetChecked(false)
		}
	}

	if par.steps/par.beats > 1 && w.fillCheckbox.Checked {
		fillSteps(w, &pattern)
	}
	par.pattern = &pattern
	return nil
}
