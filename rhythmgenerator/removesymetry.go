// RemoveSymmetry takes patterns that steps%beats != 0
// it asseses whether they are symetrical and if yes
// it finds the repeated units and reverses the last of them.
// RemoveSymmetry only works on uninverted patterns that begin with
// an 'X'.

package rhythmgenerator

import (
	"reflect"
)

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67,
	71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163,
	167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263,
	269, 271, 277, 281, 283, 293}

func removeSymmetry(steps, beats int, pattern string) string {
	if pattern == "" {
		return ""
	}
	if isPrime := isPrime(steps); steps%beats == 0 || isPrime || steps < 9 || beats < 3 {
		return ""
	}

	symmetryAxis := findAxis(steps, pattern)
	if symmetryAxis == 0 {
		return ""
	}
	pattern = swapLast(steps, pattern, symmetryAxis)
	return pattern
}

// Helpers to remove symmetry
func findAxis(steps int, pattern string) int {
	var period int
	for _, prime := range primes {
		if prime >= steps/2 {
			return 0
		}
		if steps%prime == 0 {
			period = steps / prime
			// checking if all periods contain the same motif
			for j := period; j <= steps-period; j += period {
				if pattern[:period] != pattern[j:j+period] {
					return 0
				}
			}
			return period
		}
	}
	return 0
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
			return false
		}
		if n == i {
			return true
		}
	}
	return false
}
