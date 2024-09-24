// newPattern adds a lower case 'x' on every "Xoo" block
// if there is a pattern of "Xooo" or more 'o's it calls
// Euclidean on this block and all 'x's but the first
// a lower case 'x'

package rhythmgenerator

func fillSteps(w *widgets, pattern *string) {
	var finalPattern string
	var count int
	var filledSteps bool
	parts := []int{}
	newPattern := *pattern

	if w.inversionStatus != 0 {
		newPattern = unInvertPattern(newPattern, w)
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

		if j == 2 {
			finalPattern += string(onSet) + string(offSet)
		} else if j == 3 {
			finalPattern += string(onSet) + string(offSet) + string(fill)
			filledSteps = true
		} else {
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

	if w.inversionStatus != 0 && filledSteps {
		finalPattern = reInvertPattern(finalPattern, w)
	}

	if filledSteps {
		w.genPattern.SetText(finalPattern)
		*pattern = finalPattern
		filledButtonState(w, true)
	} else {
		filledButtonState(w, false)
	}
}

func undofillSteps(w *widgets, pattern *string) {
	var filledPattern string = *pattern
	var newPattern string
	for _, char := range filledPattern {
		if char == fill {
			newPattern += string(offSet)
		} else {
			newPattern += string(char)
		}
	}
	w.genPattern.SetText(newPattern)
	*pattern = newPattern
}
