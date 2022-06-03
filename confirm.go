package fyneselfupdate

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"golang.org/x/net/context"
)

// NewConfirmCallback creates a GUI based confirmation to use as selfupdate callback
func NewConfirmCallback(win fyne.Window) func(string) bool {
	return NewConfirmCallbackWithTimeout(win, 0)
}

// NewConfirmCallbackWithTimeout creates a GUI based confirmation callback
// with a timeout, after which time the question will be confirmed.
// This can assist in a "default to yes" update where computer may be unattended.
func NewConfirmCallbackWithTimeout(win fyne.Window, timeout time.Duration) func(string) bool {
	return func(info string) bool {
		var cancel func()
		var d dialog.Dialog
		resp := make(chan bool)
		if timeout > 0 {
			ctx, fn := context.WithTimeout(context.Background(), timeout)
			cancel = fn
			go func() {
				<-ctx.Done()
				d.Hide()
				resp <- true
			}()
		}
		d = dialog.NewConfirm("Application Update", info+"\n\nDo you wish to update?\n", func(ok bool) {
			if cancel != nil {
				cancel()
			}
			resp <- ok
		}, win)

		d.Show()
		return <-resp
	}
}

// NewRestartConfirmCallback create a GUI based confirmation to approve restarting the application after being updated
func NewRestartConfirmCallback(win fyne.Window) func() bool {
	return NewRestartConfirmCallbackWithTimeout(win, false, 0)
}

// NewRestartConfirmCallbackWithTimeout creates a GUI based restarting confirmation callback
// with a timeout, after which time the question will be confirmed if defaultRestart is true.
// This can assist in a "default to yes" update where computer may be unattended.
func NewRestartConfirmCallbackWithTimeout(win fyne.Window, defaultRestart bool, timeout time.Duration) func() bool {
	return func() bool {
		var cancel func()
		var d dialog.Dialog
		resp := make(chan bool)
		if timeout > 0 {
			ctx, fn := context.WithTimeout(context.Background(), timeout)
			cancel = fn
			go func() {
				<-ctx.Done()
				d.Hide()
				resp <- defaultRestart
			}()
		}
		d = dialog.NewConfirm("Application Update", "Do you wish to restart this application?\n", func(ok bool) {
			if cancel != nil {
				cancel()
			}
			resp <- ok
		}, win)

		d.Show()
		return <-resp
	}
}
