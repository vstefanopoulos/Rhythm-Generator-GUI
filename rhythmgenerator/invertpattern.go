package rhythmgenerator

import "fmt"

// invertRight and invertLeft invert the pattern string by one step
// and update the pattern pointer and label
func invertRight(pattern *string, w *Widgets, p *Parameters) {
	invertedPattern := *pattern
	invertedPattern = invertedPattern[len(invertedPattern)-1:] + invertedPattern[0:len(invertedPattern)-1]
	p.inversionDegree = (p.inversionDegree + 1) % len(invertedPattern)
	w.inversionLabel.SetText(fmt.Sprintf("Inversion Status: %v", p.inversionDegree))
	*pattern = invertedPattern
	w.updatePatternLabel(*p.pattern)
}

func invertLeft(pattern *string, w *Widgets, p *Parameters) {
	invertedPattern := *pattern
	invertedPattern = invertedPattern[1:] + invertedPattern[0:1]
	p.inversionDegree = (p.inversionDegree - 1) % len(invertedPattern)
	w.inversionLabel.SetText(fmt.Sprintf("Inversion Status: %v", p.inversionDegree))
	*pattern = invertedPattern
	w.updatePatternLabel(*p.pattern)
}

// These are helper functions to fill steps and remove symmetry
func unInvert(pattern string, p *Parameters) string {
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

func reInvert(pattern string, p *Parameters) string {
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
