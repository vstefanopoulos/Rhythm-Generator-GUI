// newPattern adds a lower case 'x' on every "Xoo" block
// if there is a pattern of "Xooo" or more 'o's it calls
// Euclidean on this block and all 'x's but the first
// a lower case 'x'

package rhythmgenerator

func fillSteps(pattern string) (string, bool) {
	var newPattern string
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

	for _, j := range parts {

		if j == 2 {
			newPattern += string(onSet) + string(offSet)
		} else if j == 3 {
			newPattern += string(onSet) + string(offSet) + string(fill)
			filledSteps = true
		} else {
			part := euclideanGenerate(j, j/2)
			for i, char := range part {
				if i != 0 && char == 'X' {
					newPattern += "x"
				} else {
					newPattern += string(char)
				}
			}
			filledSteps = true
		}
	}
	return newPattern, filledSteps
}
