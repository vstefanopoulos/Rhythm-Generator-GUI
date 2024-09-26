// newPattern adds a lower case 'x' on every "Xoo" block
// if there is a pattern of "Xooo" or more 'o's it calls
// Euclidean on this block and all 'x's but the first
// a lower case 'x'

package rhythmgenerator

func fillSteps(w *Widgets, p *Parameters, pattern *string) {
	if *p.pattern == "" {
		return
	} else if !(p.steps/p.beats > 1) {
		filledOk(w, false)
		return
	}

	var finalPattern string
	var count int
	var filledSteps bool
	parts := []int{}
	newPattern := *pattern

	if p.inversionDegree != 0 {
		newPattern = unInvertPattern(newPattern, p)
	}

	for i, j := range newPattern {
		if j == onSet && i != 0 {
			parts = append(parts, i-count)
			count = i
		}

		if i == len(newPattern)-1 {
			parts = append(parts, i+1-count)
		}
	}

	for _, j := range parts {
		switch {
		case j == 1:
			finalPattern += string(onSet)
		case j == 2:
			finalPattern += string(onSet) + string(offSet)
		case j == 3:
			finalPattern += string(onSet) + string(offSet) + string(fill)
			filledSteps = true
		case j > 3:
			part := euclideanGenerate(j, j/2)
			for i, char := range part {
				if i != 0 && char == onSet {
					finalPattern += string(fill)
				} else {
					finalPattern += string(char)
				}
			}
			filledSteps = true
		}
	}

	if p.inversionDegree != 0 && filledSteps {
		finalPattern = reInvertPattern(finalPattern, p)
	}

	if filledSteps {
		w.updatePatternLabel(finalPattern)
		*pattern = finalPattern
		filledOk(w, true)
	} else {
		filledOk(w, false)
	}
}

func undofillSteps(w *Widgets, pattern *string) {
	if *pattern == "" {
		return
	}
	var filledPattern string = *pattern
	var newPattern string
	for _, char := range filledPattern {
		if char == fill {
			newPattern += string(offSet)
		} else {
			newPattern += string(char)
		}
	}
	w.updatePatternLabel(newPattern)
	w.patternLabel.SetText(newPattern)
	*pattern = newPattern
}
