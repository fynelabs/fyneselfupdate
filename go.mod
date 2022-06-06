module github.com/fynelabs/fyneselfupdate

go 1.14

require (
	fyne.io/fyne/v2 v2.1.4
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
)

require github.com/fynelabs/selfupdate v0.0.0-20220606220549-2086dc62c04a

replace github.com/fynelabs/selfupdate => github.com/fynelabs/selfupdate v0.0.0-20220606223111-6c71e1845abc
