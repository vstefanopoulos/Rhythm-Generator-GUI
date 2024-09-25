package rhythmgenerator

type Error struct {
	Message  string
	Solution string
}

func (err *Error) handleInputErrors(w *Widgets) {
	if err != nil {
		w.inversionStatusLabel.SetText(err.Message)
		if err.Solution != "" {
			w.genPattern.SetText(err.Solution)
		}
		initialButtonState(w)
	}
}
