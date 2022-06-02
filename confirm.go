package fyneselfupdate

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"golang.org/x/net/context"
)

func NewConfirmCallback(win fyne.Window) func(string) bool {
	return NewConfirmCallbackWithTimeout(win, 0)
}

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
		d = dialog.NewConfirm("Application Update", info + "\n\nDo you wish to update?\n", func(ok bool) {
			if cancel != nil {
				cancel()
			}
			resp <- ok
		}, win)

		d.Show()
		return <- resp
	}
}

