package rhythmgenerator

import "fmt"

func invertPattern(pattern string, w *widgets, right bool) string {
	if right {
		pattern = pattern[len(pattern)-1:] + pattern[0:len(pattern)-1]
		w.genPattern.SetText(pattern)
		w.inversionStatus = (w.inversionStatus + 1) % len(pattern)
	} else {
		pattern = pattern[1:] + pattern[0:1]
		w.genPattern.SetText(pattern)
		w.inversionStatus = (w.inversionStatus - 1) % len(pattern)
	}
	w.inversionStatusLabel.SetText(fmt.Sprintf("Inversion Status: %v", w.inversionStatus))
	return pattern
}
