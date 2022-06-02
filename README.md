# Fyne Selfupdate
An extension of the `selfupdate` repository that handles certain Fyne integrations.

## Callbacks

You can use the confirm and progress callbacks provided here to present graphical
status and request in your Fyne app.

```go
func NewConfirmCallback(fyne.Window)

func NewConfirmCallbackWithTimeout(fyne.Window, time.Duration)

func NewProgressCallback(win fyne.Window)
```
