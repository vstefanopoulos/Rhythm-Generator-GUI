// The customGenerate splits the initial pattern into groups of len = beats
// it takes the first of every group and puts it one after another
// then it takes the last of each group and puts it one after another
// and it carries on like that until there are no more beats

package rhythmgenerator

func customGenerate(steps, beats int) string {
	var index int
	var pattern string
	initialPattern := initiatePattern(steps, beats)
	left := 0
	right := beats - 1

	for left <= right {
		index = left
		for j := 0; j+index < steps; j += beats {
			pattern += string(initialPattern[index+j])
		}

		if left != right {
			index = right
			for j := 0; j+index < steps; j += beats {
				pattern += string(initialPattern[index+j])
			}
		}
		left++
		right--
	}
	return pattern
}
