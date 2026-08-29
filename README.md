# mdi

[Material Design Icons](https://pictogrammers.com/library/mdi/) (7447 SVGs, v7.4.47) as [Fyne](https://fyne.io) resources.

Every icon is its own function backed by its own `//go:embed` variable, so the
linker dead-code-eliminates the icons you don't call — your binary only
contains the SVGs you actually use.

```sh
go get github.com/roffe/mdi
```

## Usage

```go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/roffe/mdi"
)

func main() {
	a := app.New()
	w := a.NewWindow("mdi demo")

	save := theme.NewThemedResource(mdi.IconContentSave()) // follows the theme's foreground color
	w.SetContent(widget.NewButtonWithIcon("Save", save, func() {}))
	w.ShowAndRun()
}
```

Each icon function returns the raw SVG resource (black fill); wrap it in
`theme.NewThemedResource` to have it follow the theme.

```go
// Toolbar
toolbar := widget.NewToolbar(
	widget.NewToolbarAction(theme.NewThemedResource(mdi.IconPlus()), addFunc),
	widget.NewToolbarAction(theme.NewThemedResource(mdi.IconDelete()), deleteFunc),
	widget.NewToolbarSeparator(),
	widget.NewToolbarAction(theme.NewThemedResource(mdi.IconCog()), settingsFunc),
)

// Plain icon widget
icon := widget.NewIcon(theme.NewThemedResource(mdi.IconHome()))
```

## All icons / lookup by name

Dynamic lookup by string defeats dead-code elimination, so it lives in the
separate `all` subpackage. Importing it pins **all 7447 icons** into your
binary (~9 MB):

```go
import "github.com/roffe/mdi/all"

r := all.Icons["chevron-down"]() // name → fyne.Resource

for _, name := range all.Names() { // sorted names of every icon
	fmt.Println(name)
}
```

## Icon picker

An icon browser lives in [example/](example/):

```sh
go run ./example
```

![Icon picker](mdi.jpg)

Type in the search box to filter all 7447 icons live; click an icon to copy
its name to the clipboard.

## Updating

```sh
./update.sh
```

Checks npm for a newer [`@mdi/svg`](https://www.npmjs.com/package/@mdi/svg), re-downloads the SVGs, regenerates `icons.go` and `all/all.go`, and runs the tests. No-op if already up to date.

## License

Icons: [Pictogrammers Free License / Apache 2.0](LICENSE). Go code: same repo license.
