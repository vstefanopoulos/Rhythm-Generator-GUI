package rhythmgenerator

import "fmt"

func invertPattern(pattern *string, w *widgets, right bool) {
	newPattern := *pattern
	if right {
		newPattern = newPattern[len(newPattern)-1:] + newPattern[0:len(newPattern)-1]
		w.inversionStatus = (w.inversionStatus + 1) % len(*pattern)
	} else {
		newPattern = newPattern[1:] + newPattern[0:1]
		w.inversionStatus = (w.inversionStatus - 1) % len(newPattern)
	}
	w.inversionStatusLabel.SetText(fmt.Sprintf("Inversion Status: %v", w.inversionStatus))
	w.genPattern.SetText(newPattern)
	*pattern = newPattern
}
