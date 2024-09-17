package rhythmgenerator

import (
	"strconv"
)

func initiatePattern(steps, beats int) [][]rune {
	var initialPattern [][]rune
	var entry []rune

	for i := 0; i < steps; i++ {
		if i < beats {
			entry = []rune{onSet}
		} else {
			entry = []rune{offSet}
		}
		initialPattern = append(initialPattern, entry)
	}
	InputError = ""
	return initialPattern
}

func checkInput(args []string) (int, int, int) {

	steps, err := strconv.Atoi(args[0])
	if err != nil || steps < 1 {
		InputError = "Invalid Steps Input"
		InputErrorSolution = "Steps must be a possitive integer"
		return 0, 0, 0
	}

	beats, err := strconv.Atoi(args[1])
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

	bpm, err := strconv.Atoi(args[2])
	if err != nil || bpm > 300 || bpm < 1 {
		InputError = ("Invalid BPM input")
		InputErrorSolution = ("BPM must be an integer between 1-300")
		return 0, 0, 0
	}
	bpm *= 2

	return steps, beats, bpm
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
