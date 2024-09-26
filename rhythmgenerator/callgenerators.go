package rhythmgenerator

const onSet = 'X'
const offSet = 'o'
const fill = 'x'

func callGenerators(w *Widgets, p *Parameters) *Error {
	var err *Error
	p.steps, p.beats, p.bpm, err = convertInput(w)

	if err != nil {
		return err
	}

	p.euclidean = euclideanGenerate(p.steps, p.beats)
	p.custom = customGenerate(p.steps, p.beats)

	if w.algorithmTypeCheck.Checked {
		*p.pattern = p.custom
	} else {
		*p.pattern = p.euclidean
	}

	if w.removeSymmetryCheck.Checked {
		removeSymmetry(w, *p.pattern, p)
	}

	if w.fillCheck.Checked {
		fillSteps(w, p, p.pattern)
	}

	if p.inversionDegree != 0 {
		*p.pattern = reInvert(*p.pattern, p)
	}
	return nil
}

func chooseAlgorithm(w *Widgets, p *Parameters, t bool) {
	var pattern string

	if t {
		pattern = p.custom
	} else {
		pattern = p.euclidean
	}

	pattern = reInvert(pattern, p)

	if w.removeSymmetryCheck.Checked {
		removeSymmetry(w, *p.pattern, p)
	}

	if w.fillCheck.Checked {
		fillSteps(w, p, &pattern)
	}

	*p.pattern = pattern
	w.updatePatternLabel(*p.pattern)
}
