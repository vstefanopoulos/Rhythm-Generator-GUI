package rhythmgenerator

import "strconv"

type Error struct {
	Message  string
	Solution string
}

func convertInput(w *Widgets, p *Parameters) *Error {
	e := &Error{}
	var err error
	p.steps, err = strconv.Atoi(w.stepsInput.Text)
	if err != nil {
		e.Message = "Invalid Steps Input"
		e.Solution = err.Error()
		return e
	}

	if p.steps < 1 {
		e.Message = "Invalid Steps Input"
		e.Solution = "Steps must be a possitive integer"
		return e
	}

	p.beats, err = strconv.Atoi(w.beatsInput.Text)
	if err != nil {
		e.Message = "Invalid Beats Input"
		e.Solution = err.Error()
		return e
	}

	if p.beats < 1 {
		e.Message = "Invalid Beats Input"
		e.Solution = "Beat must be a possitive integer"
		return e
	}

	if p.beats > p.steps {
		e.Message = "Invalid Steps Input"
		e.Solution = "Beats cannot be more than steps"
		return e
	}

	p.bpm, err = strconv.Atoi(w.bpmInput.Text)
	if err != nil {
		e.Message = ("Invalid BPM input")
		e.Solution = err.Error()
		return e
	}

	if p.bpm > 500 {
		e.Message = ("Too fast for me")
		e.Solution = ("Fastest I can go is 500bpm")
		return e
	} else if p.bpm < 1 {
		e.Message = ("Too slow")
		e.Solution = ("BPM value should be over 0")
		return e
	}

	if w.doubletimeCheck.Checked {
		if p.bpm > 1000 {
			e.Message = ("Too fast for me")
			e.Solution = ("When double time the BPM limit is 250")
			return e
		}
	}
	return nil
}
