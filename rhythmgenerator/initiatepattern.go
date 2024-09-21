package rhythmgenerator

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
	return initialPattern
}
