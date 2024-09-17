package rhythmgenerator

func euclideanGenerate(steps, beats int) string {
	var result string
	pattern := initiatePattern(steps, beats)
	d := steps - beats
	steps, beats = max(d, beats), min(d, beats)
	z := d

	for z > 0 || d > 1 {

		for i := 0; i < beats; i++ {
			pattern[i] = append(pattern[i], pattern[len(pattern)-1-i]...)
		}
		pattern = pattern[:len(pattern)-beats]
		z -= beats
		d = steps - beats
		steps, beats = max(d, beats), min(d, beats)
	}

	for i := range pattern {

		for _, k := range pattern[i] {
			result += string(k)
		}
	}
	return result
}
