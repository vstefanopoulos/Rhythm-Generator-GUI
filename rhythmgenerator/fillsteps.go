// FillSteps adds a lower case 'x' on every "Xoo" block
// if there is a pattern of "Xooo" or more 'o's it calls
// Euclidean on this block and makes all 'x's but the first
// a lower case 'x'. FillSteps only works on uninverted patterns
// that begin with an 'X'.

package rhythmgenerator

func fillSteps(steps, beats int, pattern string) string {
	if pattern == "" {
		return ""
	} else if !(steps/beats > 1) {
		return ""
	}

	var count int
	var filledSteps bool
	parts := []int{}

	for i, j := range pattern {
		if j == onSet && i != 0 {
			parts = append(parts, i-count)
			count = i
		}

		if i == len(pattern)-1 {
			parts = append(parts, i+1-count)
		}
	}
	var filledPattern string
	for _, j := range parts {
		switch {
		case j == 1:
			filledPattern += string(onSet)
		case j == 2:
			filledPattern += string(onSet) + string(offSet)
		case j == 3:
			filledPattern += string(onSet) + string(offSet) + string(fill)
			filledSteps = true
		case j > 3:
			part := euclideanGenerate(j, j/2)
			for i, char := range part {
				if i != 0 && char == onSet {
					filledPattern += string(fill)
				} else {
					filledPattern += string(char)
				}
			}
			filledSteps = true
		}
	}

	if filledSteps {
		return filledPattern
	} else {
		return ""
	}
}
