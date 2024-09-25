package rhythmgenerator

import "fmt"

func invertPattern(pattern *string, w *Widgets, p *Parameters, right bool) {
	newPattern := *pattern
	if right {
		newPattern = newPattern[len(newPattern)-1:] + newPattern[0:len(newPattern)-1]
		p.inversionStatus = (p.inversionStatus + 1) % len(*pattern)
	} else {
		newPattern = newPattern[1:] + newPattern[0:1]
		p.inversionStatus = (p.inversionStatus - 1) % len(newPattern)
	}
	w.inversionStatusLabel.SetText(fmt.Sprintf("Inversion Status: %v", p.inversionStatus))
	w.patternLabel.SetText(newPattern)
	*pattern = newPattern
}

func unInvertPattern(pattern string, p *Parameters) string {
	switch {
	case p.inversionStatus > 0:
		index := p.inversionStatus
		pattern = pattern[index:] + pattern[:index]
	case p.inversionStatus < 0:
		index := len(pattern) + p.inversionStatus
		pattern = pattern[index:] + pattern[:index]
	}
	return pattern
}

func reInvertPattern(pattern string, p *Parameters) string {
	switch {
	case p.inversionStatus > 0:
		index := len(pattern) - p.inversionStatus
		pattern = pattern[index:] + pattern[:index]
	case p.inversionStatus < 0:
		index := -p.inversionStatus
		pattern = pattern[index:] + pattern[:index]
	}
	return pattern
}
