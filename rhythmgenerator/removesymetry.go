// RemoveSymetry takes patterns that steps%beats != 0
// it asseses wheteher they are symetrical and if yes
// it finds the repearing unit and reverses the last of them

package rhythmgenerator

import "reflect"

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

func removeSymetry(pattern string, steps int) (string, bool) {
	isPrime := isPrime(steps)

	if isPrime {
		return "", false
	}

	var symetry int
	for _, i := range primes {

		if steps%i == 0 {
			period := steps / i

			for j := period; j <= steps-period; j += period {

				if pattern[:period] != pattern[j:j+period] {
					return "", false
				}
			}
			symetry = i
			break
		}
	}

	cell := string(pattern[:steps/symetry])
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
	return newPattern, true
}

// if pattern len is not prime and not 1 then it is probably symetrical
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
