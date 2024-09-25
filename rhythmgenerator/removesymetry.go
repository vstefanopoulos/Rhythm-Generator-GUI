// RemoveSymetry takes patterns that steps%beats != 0
// it asseses wheteher they are symetrical and if yes
// it finds the repearing unit and reverses the last of them

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
		pattern = unInvertPattern(pattern, p)
	}

	symmetryAxis := findAxis(steps, pattern)
	if symmetryAxis == 0 {
		rsOk(w, false)
		return
	}

	cell := string(pattern[:steps/symmetryAxis])
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

	newPattern := pattern[:steps-len(cell)]
	for i, j := range sets {
		for k := range j {
			newPattern += string(sets[i][k])
		}
	}

	if p.inversionDegree != 0 {
		newPattern = reInvertPattern(newPattern, p)
	}
	if w.fillCheck.Checked {
		fillSteps(w, p, &newPattern)
	}

	p.pattern = &newPattern
	w.updatePatternLabel(*p.pattern)
	rsOk(w, true)
}

func findAxis(steps int, pattern string) int {
	var symmetryAxis int
	for _, i := range primes {
		if steps%i == 0 {
			period := steps / i
			for j := period; j <= steps-period; j += period {
				if pattern[:period] != pattern[j:j+period] {
					return 0
				}
			}
			symmetryAxis = i
			break
		}
	}
	return symmetryAxis
}

func fallBack(w *Widgets, p *Parameters) {
	var fallBackPattern string

	if w.algorithmTypeCheck.Checked {
		fallBackPattern = customGenerate(p.steps, p.beats)
	} else {
		fallBackPattern = euclideanGenerate(p.steps, p.beats)
	}

	if p.inversionDegree != 0 {
		fallBackPattern = reInvertPattern(fallBackPattern, p)
	}

	if w.fillCheck.Checked {
		fillSteps(w, p, &fallBackPattern)
	}
	p.pattern = &fallBackPattern
	w.updatePatternLabel(*p.pattern)
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
