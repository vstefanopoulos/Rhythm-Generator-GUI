package rhythmgenerator

import "fmt"

func invertPattern(pattern *string, w *Widgets, p *Parameters, right bool) {
	newPattern := *pattern
	if right {
		newPattern = newPattern[len(newPattern)-1:] + newPattern[0:len(newPattern)-1]
		p.inversionDegree = (p.inversionDegree + 1) % len(newPattern)
	} else {
		newPattern = newPattern[1:] + newPattern[0:1]
		p.inversionDegree = (p.inversionDegree - 1) % len(newPattern)
	}

	w.inversionLabel.SetText(fmt.Sprintf("Inversion Status: %v", p.inversionDegree))
	*pattern = newPattern
	w.updatePatternLabel(*p.pattern)
}

func unInvertPattern(pattern string, p *Parameters) string {
	switch {
	case p.inversionDegree > 0:
		index := p.inversionDegree
		pattern = pattern[index:] + pattern[:index]
	case p.inversionDegree < 0:
		index := len(pattern) + p.inversionDegree
		pattern = pattern[index:] + pattern[:index]
	}
	return pattern
}

func reInvertPattern(pattern string, p *Parameters) string {
	switch {
	case p.inversionDegree > 0:
		index := len(pattern) - p.inversionDegree
		pattern = pattern[index:] + pattern[:index]
	case p.inversionDegree < 0:
		index := -p.inversionDegree
		pattern = pattern[index:] + pattern[:index]
	}
	return pattern
}
