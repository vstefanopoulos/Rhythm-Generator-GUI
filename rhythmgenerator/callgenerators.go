package rhythmgenerator

const onSet = 'X'
const offSet = 'o'
const fill = 'x'

func callGenerators(w *widgets, par *par) *Error {
	var err *Error
	par.steps, par.beats, par.bpm, err = convertInput(w)

	if err != nil {
		return err
	}
	w.inversionStatus = 0

	if w.algCheckbox.Checked {
		*par.pattern = customGenerate(par.steps, par.beats)
	} else {
		*par.pattern = euclideanGenerate(par.steps, par.beats)
	}

	if w.removeSymmetryCheckbox.Checked {
		removeSymmetry(w, *par.pattern, par)
	}

	if par.steps/par.beats > 1 && w.fillCheckbox.Checked {
		fillSteps(w, par.pattern)
	}
	return nil
}
