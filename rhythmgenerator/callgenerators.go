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
	if w.algCheckbox.Checked {
		*p.pattern = p.custom
	} else {
		*p.pattern = p.euclidean
	}

	if w.removeSymmetryCheckbox.Checked {
		removeSymmetry(w, *p.pattern, p)
	}

	if p.steps/p.beats > 1 && w.fillCheckbox.Checked {
		fillSteps(w, p, p.pattern)
	}
	return nil
}

func chooseCustom(w *Widgets, p *Parameters) {
	pattern := p.custom
	pattern = reInvertPattern(pattern, p)
	if w.removeSymmetryCheckbox.Checked {
		removeSymmetry(w, *p.pattern, p)
	}
	if w.fillCheckbox.Checked {
		fillSteps(w, p, &pattern)
	}
	w.genPattern.SetText(pattern)
	*p.pattern = pattern
}

func chooseEuclidean(w *Widgets, p *Parameters) {
	pattern := p.euclidean
	pattern = reInvertPattern(pattern, p)
	if w.removeSymmetryCheckbox.Checked {
		removeSymmetry(w, pattern, p)
	}
	if w.fillCheckbox.Checked {
		fillSteps(w, p, &pattern)
	}
	w.genPattern.SetText(pattern)
	*p.pattern = pattern
}
