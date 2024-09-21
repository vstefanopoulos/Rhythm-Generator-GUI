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
	bpm *= 2
	if bpm > 600 {
		e.Message = ("Too fast for me")
		e.Solution = ("Fastest I can go is 300bpm")
		return 0, 0, 0, &e
	} else if bpm < 1 {
		e.Message = ("Too slow")
		e.Solution = ("BPM value should be over 0")
		return 0, 0, 0, &e
	}
	if w.doubletimeCheckbox.Checked {
		bpm *= 2
		if bpm > 600 {
			e.Message = ("Too fast for me")
			e.Solution = ("Try slower BPM or not double time")
			return 0, 0, 0, &e
		}
	}
	return steps, beats, bpm, nil
}
