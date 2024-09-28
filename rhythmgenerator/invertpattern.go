package rhythmgenerator

// invertRight and invertLeft invert the pattern string by one step
// and update the pattern pointer and label
func invertRight(pattern *string, w *Widgets, p *Parameters) {
	var invertedPattern string
	if *pattern == "" {
		p.inversionDegree = (p.inversionDegree + 1)
		w.updateInversionLabel(p.inversionDegree)
		return
	}
	invertedPattern = *pattern
	invertedPattern = invertedPattern[len(invertedPattern)-1:] + invertedPattern[0:len(invertedPattern)-1]
	*pattern = invertedPattern
	w.updatePatternLabel(*p.pattern)
	p.inversionDegree = (p.inversionDegree + 1) % len(invertedPattern)
	w.updateInversionLabel(p.inversionDegree)
}

func invertLeft(pattern *string, w *Widgets, p *Parameters) {
	var invertedPattern string
	if *pattern == "" {
		p.inversionDegree = (p.inversionDegree - 1)
		w.updateInversionLabel(p.inversionDegree)
		return
	}
	invertedPattern = *pattern
	invertedPattern = invertedPattern[1:] + invertedPattern[0:1]
	*pattern = invertedPattern
	w.updatePatternLabel(*p.pattern)
	p.inversionDegree = (p.inversionDegree - 1) % len(invertedPattern)
	w.updateInversionLabel(p.inversionDegree)
}

func invertByDegree(pattern string, p *Parameters) string {
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
