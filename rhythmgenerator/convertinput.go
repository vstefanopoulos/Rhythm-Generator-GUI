package rhythmgenerator

import "strconv"

func convertInput(w *widgets) (int, int, int, *Error) {
	e := Error{}
	steps, err := strconv.Atoi(w.stepsInput.Text)
	if err != nil {
		e.Message = "Invalid Steps Input"
		e.Solution = err.Error()
		return 0, 0, 0, &e
	}
	if steps < 1 {
		e.Message = "Invalid Steps Input"
		e.Solution = "Steps must be a possitive integer"
		return 0, 0, 0, &e
	}

	beats, err := strconv.Atoi(w.beatsInput.Text)
	if err != nil {
		e.Message = "Invalid Beats Input"
		e.Solution = err.Error()
		return 0, 0, 0, &e
	}
	if beats < 1 {
		e.Message = "Invalid Beats Input"
		e.Solution = "Beat must be a possitive integer"
		return 0, 0, 0, &e
	}
	if beats > steps {
		e.Message = "Invalid Steps Input"
		e.Solution = "Beats cannot be more than steps"
		return 0, 0, 0, &e
	}

	bpm, err := strconv.Atoi(w.bpmInput.Text)
	if err != nil {
		e.Message = ("Invalid BPM input")
		e.Solution = err.Error()
		return 0, 0, 0, &e
	}
	if bpm > 300 || bpm < 1 {
		e.Message = ("Invalid BPM input")
		e.Solution = ("BPM must be an integer between 1-300")
		return 0, 0, 0, &e
	}
	bpm *= 2

	return steps, beats, bpm, nil
}
