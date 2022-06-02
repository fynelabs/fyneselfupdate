package fyneselfupdate

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// NewProgressCallback returns a callback that can be used to present download information during update.
func NewProgressCallback(win fyne.Window) func(float64, error) {
	var d dialog.Dialog
	return func(done float64, fail error) {
		if d == nil {
			if done < 0 {
				d = dialog.NewProgressInfinite("Application update", "Downloading update", win)
			} else {
				d = dialog.NewProgress("Application update", "Downloading update", win)
			}
			d.Show()
		}

		if fail != nil {
			d.Hide()
			d = nil
			dialog.ShowError(fail, win)
			return
		}
		if done == 1.0 {
			d.Hide()
			d = nil
			return
		}

		if progress, ok := d.(*dialog.ProgressDialog); ok {
			progress.SetValue(done)
		}
	}
}
