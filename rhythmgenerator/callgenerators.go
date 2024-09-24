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
	par.euclidean = euclideanGenerate(par.steps, par.beats)
	par.custom = customGenerate(par.steps, par.beats)
	if w.algCheckbox.Checked {
		*par.pattern = par.custom
	} else {
		*par.pattern = par.euclidean
	}

	if w.removeSymmetryCheckbox.Checked {
		removeSymmetry(w, *par.pattern, par)
	}

	if par.steps/par.beats > 1 && w.fillCheckbox.Checked {
		fillSteps(w, par.pattern)
	}
	return nil
}

func chooseCustom(w *widgets, par *par) {
	pattern := par.custom
	pattern = reInvertPattern(pattern, w)
	if w.removeSymmetryCheckbox.Checked {
		removeSymmetry(w, *par.pattern, par)
	}
	if w.fillCheckbox.Checked {
		fillSteps(w, &pattern)
	}
	w.genPattern.SetText(pattern)
	*par.pattern = pattern
}

func chooseEuclidean(w *widgets, par *par) {
	pattern := par.euclidean
	pattern = reInvertPattern(pattern, w)
	if w.removeSymmetryCheckbox.Checked {
		removeSymmetry(w, pattern, par)
	}
	if w.fillCheckbox.Checked {
		fillSteps(w, &pattern)
	}
	w.genPattern.SetText(pattern)
	*par.pattern = pattern
}
