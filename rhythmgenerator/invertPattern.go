package rhythmgenerator

import "fyne.io/fyne/v2/widget"

func invertPattern(pattern string, inverted *bool, genPattern *widget.Label, right bool) string {
	if pattern == "" {
		return pattern
	}
	if right {
		pattern = pattern[len(pattern)-1:] + pattern[0:len(pattern)-1]
		genPattern.SetText(pattern)
		*inverted = true
	} else {
		pattern = pattern[1:] + pattern[0:1]
		genPattern.SetText(pattern)
		*inverted = true
	}
	return pattern
}
