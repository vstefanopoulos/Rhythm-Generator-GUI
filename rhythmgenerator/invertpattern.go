package rhythmgenerator

import "fmt"

func invertPattern(pattern *string, w *widgets, right bool) {
	oldPattern := *pattern
	var newPattern string
	if right {
		newPattern = oldPattern[len(oldPattern)-1:] + oldPattern[0:len(oldPattern)-1]
		w.inversionStatus = (w.inversionStatus + 1) % len(*pattern)
	} else {
		newPattern = oldPattern[1:] + oldPattern[0:1]
		w.inversionStatus = (w.inversionStatus - 1) % len(oldPattern)
	}
	w.inversionStatusLabel.SetText(fmt.Sprintf("Inversion Status: %v", w.inversionStatus))
	w.genPattern.SetText(newPattern)
	*pattern = newPattern
}
