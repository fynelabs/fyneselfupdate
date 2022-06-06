package main

import (
	"crypto/ed25519"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/fynelabs/fyneselfupdate"
	"github.com/fynelabs/selfupdate"
)

func selfManage(a fyne.App, w fyne.Window) {
	// Used `selfupdatectl create-keys` followed by `selfupdatectl print-key`
	publicKey := ed25519.PublicKey{178, 103, 83, 57, 61, 138, 18, 249, 244, 80, 163, 162, 24, 251, 190, 241, 11, 168, 179, 41, 245, 27, 166, 70, 220, 254, 118, 169, 101, 26, 199, 129}

	// The public key above match the signature of the below file served by our CDN and uploaded with `selfupdatectl upload`
	httpSource := selfupdate.NewHTTPSource(nil, "http://geoffrey-test-artefacts.fynelabs.com/nomad-{{.OS}}-{{.Arch}}{{.Ext}}")

	config := fyneselfupdate.NewConfigWithTimeout(a, w, time.Duration(1)*time.Minute,
		httpSource,
		selfupdate.Schedule{FetchOnStart: true, Interval: time.Hour * time.Duration(12)},
		publicKey)

	config.Current = &selfupdate.Version{Date: time.Unix(100, 0)}

	_, err := selfupdate.Manage(config)
	if err != nil {
		log.Println("Error while setting up update manager: ", err)
		return
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	selfManage(a, w)

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewBorder(nil, nil, nil, nil, container.NewVBox(
		layout.NewSpacer(),
		container.NewHBox(layout.NewSpacer(), hello, layout.NewSpacer()),
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
		layout.NewSpacer(),
	)))
	w.Resize(fyne.Size(fyne.NewSize(640, 480)))

	w.ShowAndRun()
}
