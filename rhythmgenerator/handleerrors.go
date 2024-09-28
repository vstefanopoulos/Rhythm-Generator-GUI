package rhythmgenerator

type Error struct {
	Message  string
	Solution string
}

func handleErrors(w *Widgets, p *Parameters) *Error {
	err := convertInput(w, p)
	if err != nil {
		w.errLabel.Text = err.Message
		if err.Solution != "" {
			w.errSolutionLabel.Text = err.Solution
		}
		return err
	}
	w.errLabel.Text = ""
	w.errLabel.Refresh()
	w.errSolutionLabel.Text = ""
	w.errSolutionLabel.Refresh()
	return nil
}
