package rhythmgenerator

import "strconv"

func convertInput(stepsInput, beatsInput, bpmInput string) (int, int, int) {

	steps, err := strconv.Atoi(stepsInput)
	if err != nil || steps < 1 {
		InputError = "Invalid Steps Input"
		InputErrorSolution = "Steps must be a possitive integer"
		return 0, 0, 0
	}

	beats, err := strconv.Atoi(beatsInput)
	if err != nil || beats < 1 {
		InputError = "Invalid Beats Input"
		InputErrorSolution = "Beat must be a possitive integer"
		return 0, 0, 0
	}
	if beats > steps {
		InputError = "Invalid Steps Input"
		InputErrorSolution = "Beats cannot be more than steps"
		return 0, 0, 0
	}

	bpm, err := strconv.Atoi(bpmInput)
	if err != nil || bpm > 300 || bpm < 1 {
		InputError = ("Invalid BPM input")
		InputErrorSolution = ("BPM must be an integer between 1-300")
		return 0, 0, 0
	}
	bpm *= 2

	return steps, beats, bpm
}
