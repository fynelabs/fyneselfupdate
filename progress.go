package fyneselfupdate

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// NewProgressCallback returns a callback that can be used to present download information during update.
func NewProgressCallback(win fyne.Window) func(float64, error) {
	var d dialog.Dialog
	var progress fyne.Widget
	return func(done float64, fail error) {
		if d == nil {
			if done < 0 {
				infinite := widget.NewProgressBarInfinite()
				infinite.Start()
				progress = infinite
			} else {
				progress = widget.NewProgressBar()
			}
			content := container.NewVBox(widget.NewLabel("Downloading update"), progress)
			d = dialog.NewCustomWithoutButtons("Application update", content, win)
			d.Show()
		}

		cleanup := func() {
			if infinite, ok := progress.(*widget.ProgressBarInfinite); ok {
				infinite.Stop()
			}

			d.Hide()
			d = nil
		}

		if fail != nil {
			dialog.ShowError(fail, win)
			cleanup()
			return
		}
		if done == 1.0 {
			cleanup()
			return
		}

		if limited, ok := progress.(*widget.ProgressBar); ok {
			limited.SetValue(done)
		}
	}
}
