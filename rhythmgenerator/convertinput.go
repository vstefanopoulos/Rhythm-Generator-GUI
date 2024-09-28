package rhythmgenerator

import "strconv"

func convertInput(w *Widgets, p *Parameters) *Error {
	customError := &Error{}
	var err error
	p.steps, err = strconv.Atoi(w.stepsInput.Text)
	if err != nil {
		customError.Message = "Invalid Steps Input"
		customError.Solution = err.Error()
		return customError
	}

	if p.steps < 1 {
		customError.Message = "Invalid Steps Input"
		customError.Solution = "Steps must be a possitive integer"
		return customError
	}

	p.beats, err = strconv.Atoi(w.beatsInput.Text)
	if err != nil {
		customError.Message = "Invalid Beats Input"
		customError.Solution = err.Error()
		return customError
	}

	if p.beats < 1 {
		customError.Message = "Invalid Beats Input"
		customError.Solution = "Beat must be a possitive integer"
		return customError
	}

	if p.beats > p.steps {
		customError.Message = "Invalid Steps Input"
		customError.Solution = "Beats cannot be more than steps"
		return customError
	}

	p.bpm, err = strconv.Atoi(w.bpmInput.Text)
	if err != nil {
		customError.Message = ("Invalid BPM input")
		customError.Solution = err.Error()
		return customError
	}

	if p.bpm > 500 {
		customError.Message = ("Too fast for me")
		customError.Solution = ("Fastest I can go is 500bpm")
		return customError
	} else if p.bpm < 1 {
		customError.Message = ("Too slow")
		customError.Solution = ("BPM value should be over 0")
		return customError
	}

	if w.doubletimeCheck.Checked {
		if p.bpm > 1000 {
			customError.Message = ("Too fast for me")
			customError.Solution = ("When double time the BPM limit is 250")
			return customError
		}
	}
	return nil
}
