package rhythmgenerator

const onSet = 'X'
const offSet = 'o'
const fill = 'x'

var InputError string
var InputErrorSolution string

func RGMain(args []string) (string, string, int) {
	var pattern string
	steps, beats, bpm := checkInput(args)
	var algType = args[3]

	// if err check input returns 0, 0
	if steps == 0 || beats == 0 {
		return "", "", 0
	}

	if algType == "Euclidean" {
		pattern = euclideanGenerate(steps, beats)
	} else {
		pattern = customGenerate(steps, beats)
	}
	if steps > 2 && beats > 3 && steps%beats != 0 {
		newPattern, isSymetrical := removeSymetry(pattern, steps)
		if isSymetrical {
			pattern = newPattern
			algType += " Asymetrical "
		}
	}

	if steps/beats > 1 && args[4] == "fill" {
		filledPattern := fillSteps(pattern)
		algType += " Filled"
		return filledPattern, algType, bpm
	}
	return pattern, algType, bpm
}
