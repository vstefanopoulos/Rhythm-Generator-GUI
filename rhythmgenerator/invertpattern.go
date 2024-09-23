package rhythmgenerator

import "fmt"

func invertPattern(pattern *string, w *widgets, right bool) {
	oldPattern := *pattern
	var newPattern string
	if right {
		newPattern = oldPattern[len(*pattern)-1:] + oldPattern[0:len(*pattern)-1]
		w.inversionStatus = (w.inversionStatus + 1) % len(*pattern)
	} else {
		newPattern = oldPattern[1:] + oldPattern[0:1]
		w.inversionStatus = (w.inversionStatus - 1) % len(*pattern)
	}
	w.inversionStatusLabel.SetText(fmt.Sprintf("Inversion Status: %v", w.inversionStatus))
	*pattern = newPattern
	w.genPattern.SetText(*pattern)
}
