package rhythmgenerator

const onSet = 'X'
const offSet = 'o'
const fill = 'x'

func callGenerators(w *Widgets, p *Parameters) {
	var pattern string

	if w.algorithmTypeCheck.Checked {
		pattern = customGenerate(p.steps, p.beats)
	} else {
		pattern = euclideanGenerate(p.steps, p.beats)
	}

	if w.removeSymmetryCheck.Checked {
		if temp := removeSymmetry(p.steps, p.beats, pattern); temp != "" {
			pattern = temp
			defer rsOk(w, true)
		} else {
			defer rsOk(w, false)
		}
	}

	if w.fillCheck.Checked {
		if temp := fillSteps(p.steps, p.beats, pattern); temp != "" {
			pattern = temp
			defer filledOk(w, true)
		} else {
			defer filledOk(w, false)
		}
	}

	if p.inversionDegree != 0 {
		pattern = invertByDegree(pattern, p)
	}

	*p.pattern = pattern
	w.updatePatternLabel(*p.pattern)
	w.updateInversionLabel(p.inversionDegree)
	return
}
