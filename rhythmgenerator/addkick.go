package rhythmgenerator

func addKick(pattern string) string {
	var newPattern string
	var kick bool = true
	for _, char := range pattern {
		if char == 'X' {
			if kick {
				char = 'O'
			}
			kick = !kick
		}
		newPattern += string(char)
	}
	return newPattern
}
