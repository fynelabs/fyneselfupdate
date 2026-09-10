# Fyne Selfupdate

[![godoc reference](https://godoc.org/github.com/fynelabs/fyneselfupdate?status.png)](https://godoc.org/github.com/fynelabs/selfupdate)
[![Coverage Status](https://coveralls.io/repos/github/fynelabs/fyneselfupdate/badge.svg?branch=main)](https://coveralls.io/github/fynelabs/fyneselfupdate?branch=main)

An extension of the `selfupdate` repository that handles certain Fyne integrations.

## Simple setup for Fyne application

If you want to add self update support to your Fyne application, you can just just add the following code
```go
	// Used `selfupdatectl create-keys` followed by `selfupdatectl print-key`
	publicKey := ed25519.PublicKey{178, 103, 83, 57, 61, 138, 18, 249, 244, 80, 163, 162, 24, 251, 190, 241, 11, 168, 179, 41, 245, 27, 166, 70, 220, 254, 118, 169, 101, 26, 199, 129}

	// The public key above match the signature of the below file served by our CDN
	httpSource := selfupdate.NewHTTPSource(nil, "http://localhost/myprogram-{{.OS}}-{{.Arch}}{{.Ext}}")

	config := fyneselfupdate.NewConfigWithTimeout(a, w, time.Duration(1)*time.Minute,
		httpSource,
		selfupdate.Schedule{FetchOnStart: true, Interval: time.Hour * time.Duration(12)},
		publicKey)

	_, err := selfupdate.Manage(config)
	if err != nil {
		log.Println("Error while setting up update manager: ", err)
		return
	}
```

For a better understanding of how to use `selfupdatectl` to help you manage your self updating Fyne application deployment you can read the documentation [here](https://github.com/fynelabs/selfupdate/tree/main/cmd/selfupdatectl)

## Callbacks

You can also use following callbacks and directly create your custom selfupdate.Config structure to adapt presented graphical element to your exact need in your Fyne app.

```go
func NewUpgradeConfirmCallback(fyne.Window)

func NewUpgradeConfirmCallbackWithTimeout(fyne.Window, time.Duration)

func NewRestartConfirmCallback(win fyne.Window)

func NewRestartConfirmCallbackWithTimeout(win fyne.Window, timeout time.Duration)

func NewProgressCallback(win fyne.Window)

func NewExitCallback(app fyne.App, win fyne.Window)
```

## Localization

`fyneselfupdate` uses Fyne's standard [`lang`](https://docs.fyne.io/api/v2/lang/pkg/) package to
translate the text shown in its dialogs (update confirmation, restart confirmation and the
download progress window). No text is hardcoded to English only, everything goes through
`lang.L`, so it will follow whatever locale your application has configured.

Ready-made translation bundles for `fyneselfupdate`'s own strings are provided in the
[`translations`](./translations) folder, named `fyneselfupdate.<lang>.json` (for example
`fyneselfupdate.de.json`). Copy the ones you need into your own application and load them
alongside your own translations:

```go
//go:embed translations/*.json
var translationsFS embed.FS

func main() {
	a := app.New()
	w := a.NewWindow("My App")

	// Load fyneselfupdate's translations together with your own app's translations.
	if err := lang.AddTranslationsFS(translationsFS, "translations"); err != nil {
		log.Println("Failed to load translations:", err)
	}

	...
}
```

The `<prefix>.` in the file name avoids collisions if your own application already has a
`de.json`, `fr.json`, etc. in the same `translations` folder, just keep both files side by side.

If a translation for the current locale is missing, `fyneselfupdate` falls back to the original
English text, so adding localization support is entirely optional and backward compatible.

Contributions of new language bundles for `fyneselfupdate` are welcome - please open a pull
request adding a `fyneselfupdate.<lang>.json` file under `translations/`.


## API Compatibility Promises
The main branch of `fyneselfupdate` is *not* guaranteed to have a stable API over time. Still we will try hard to not break its API unecessarily and will follow a proper versioning of our release. We will also keep it in sync and up to date with `fynelabs/selfupate`.

The `fyneselfupdate` package makes the following promises about API compatibility:
1. A list of all API-breaking changes will be documented in this README.
1. `fyneselfupdate` will strive for as few API-breaking changes as possible.

## License
BSD 3-Clause

## Sponsors

This project is kindly sponsored by the following companies:

<a href="https://dentagraphics.com/" style="text-decoration: none">
<img width="190" src="img/sponsor/dentagraphics.jpg" />
</a>
