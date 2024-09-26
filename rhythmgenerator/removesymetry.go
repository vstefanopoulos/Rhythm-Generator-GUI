// removeSymmetry takes patterns that steps%beats != 0
// it asseses whether they are symetrical and if yes
// it finds the repearing unit and reverses the last of them.
// RemoveSymmetry only works on un inverted and not filled patterns
// which is why filled and inversions have to be undone before and redone after.
// After the pattern generation it updates the pointer to the playing or to be played pattern
// and the pattern label

package rhythmgenerator

import (
	"reflect"
)

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

func removeSymmetry(w *Widgets, pattern string, p *Parameters) {
	if pattern == "" {
		return
	}
	steps, beats := p.steps, p.beats
	if isPrime := isPrime(steps); steps%beats == 0 || isPrime || steps < 9 || beats < 3 {
		rsOk(w, false)
		return
	}

	if w.fillCheck.Checked {
		undofillSteps(w, &pattern)
	}
	if p.inversionDegree != 0 {
		pattern = unInvert(pattern, p)
	}

	symmetryAxis := findAxis(steps, pattern)
	if symmetryAxis == 0 {
		rsOk(w, false)
		return
	}
	pattern = swapLast(steps, pattern, symmetryAxis)

	if p.inversionDegree != 0 {
		pattern = reInvert(pattern, p)
	}
	if w.fillCheck.Checked {
		fillSteps(w, p, &pattern)
	}

	p.pattern = &pattern
	w.updatePatternLabel(*p.pattern)
	rsOk(w, true)
}

// fallBack recalls the algorithms and resets the inversion and fill status.
// Then updates the pointer to the playing or to be played pattern
func fallBack(w *Widgets, p *Parameters) {
	var fallBackPattern string

	if w.algorithmTypeCheck.Checked {
		fallBackPattern = customGenerate(p.steps, p.beats)
	} else {
		fallBackPattern = euclideanGenerate(p.steps, p.beats)
	}

	if p.inversionDegree != 0 {
		fallBackPattern = reInvert(fallBackPattern, p)
	}
	if w.fillCheck.Checked {
		fillSteps(w, p, &fallBackPattern)
	}
	p.pattern = &fallBackPattern
	w.updatePatternLabel(*p.pattern)
}

// Helpers to remove symmetry
func findAxis(steps int, pattern string) int {
	var symmetryAxis int
	var period int
	for _, prime := range primes {
		if prime >= steps/2 {
			break
		}
		if steps%prime == 0 {
			period = steps / prime
			for j := period; j <= steps-period; j += period {
				if pattern[:period] != pattern[j:j+period] {
					return 0
				}
			}
			symmetryAxis = period
			break
		}
	}
	return symmetryAxis
}

func swapLast(steps int, pattern string, symmetryAxis int) string {
	cell := string(pattern[:symmetryAxis])
	var sets [][]string
	start := 0
	var end int

	for i := 0; i < len(cell); i++ {
		if cell[i] == onSet && i != 0 {
			end = i
			sets = append(sets, []string{pattern[start:end]})
			start = end
		}
		if i == len(cell)-1 {
			sets = append(sets, []string{pattern[start : i+1]})
		}
	}

	for i := len(sets) - 1; i > 0; i-- {
		if !reflect.DeepEqual(sets[i], sets[i-1]) {
			sets[i], sets[i-1] = sets[i-1], sets[i]
			break
		}
	}

	pattern = pattern[:steps-len(cell)]
	for i, j := range sets {
		for k := range j {
			pattern += string(sets[i][k])
		}
	}
	return pattern
}

func isPrime(n int) bool {
	for _, i := range primes {
		if i > n {
			break
		}
		if n == i {
			return true
		}
	}
	return false
}
