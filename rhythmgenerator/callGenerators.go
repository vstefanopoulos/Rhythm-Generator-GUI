package rhythmgenerator

const onSet = 'X'
const offSet = 'o'
const fill = 'x'

func callGenerators(stepsInput, beatsInput, bpmInput, algType string, fill, removerSymetry bool) (string, int) {
	var pattern string
	steps, beats, bpm := convertInput(stepsInput, beatsInput, bpmInput)

	// if err check input returns 0, 0
	if steps == 0 || beats == 0 {
		return "", 0
	}

	if algType == "Euclidean" {
		pattern = euclideanGenerate(steps, beats)
	} else {
		pattern = customGenerate(steps, beats)
	}
	if removerSymetry && steps > 2 && beats > 3 && steps%beats != 0 {
		newPattern, isSymetrical := removeSymetry(pattern, steps)
		if isSymetrical {
			pattern = newPattern
		}
	}
	if steps/beats > 1 && fill {
		filledPattern := fillSteps(pattern)
		algType += " Filled"
		return filledPattern, bpm
	}
	return pattern, bpm
}
