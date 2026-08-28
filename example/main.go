package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/roffe/mdi"
)

func main() {
	a := app.New()
	w := a.NewWindow("MDI icon picker")

	all := mdi.Names()
	filtered := all

	status := widget.NewLabel(fmt.Sprintf("%d icons — click one to copy its name", len(all)))

	var grid *widget.GridWrap
	grid = widget.NewGridWrap(
		func() int { return len(filtered) },
		func() fyne.CanvasObject {
			icon := widget.NewIcon(nil)
			label := widget.NewLabel("account-arrow-down") // sizes the cell
			label.Alignment = fyne.TextAlignCenter
			label.Truncation = fyne.TextTruncateEllipsis
			big := container.NewGridWrap(fyne.NewSize(48, 48), icon) // fixed 48x48 icon
			return container.NewBorder(nil, label, nil, nil, container.NewCenter(big))
		},
		func(id widget.GridWrapItemID, o fyne.CanvasObject) {
			name := filtered[id]
			box := o.(*fyne.Container)
			icon := box.Objects[0].(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*widget.Icon)
			icon.SetResource(mdi.ThemedIcon(mdi.IconName(name)))
			box.Objects[1].(*widget.Label).SetText(name)
		},
	)
	grid.OnSelected = func(id widget.GridWrapItemID) {
		name := filtered[id]
		a.Clipboard().SetContent(name)
		status.SetText(fmt.Sprintf("copied %q to clipboard", name))
		grid.UnselectAll()
	}

	search := widget.NewEntry()
	search.SetPlaceHolder("Search icons…")
	search.OnChanged = func(q string) {
		q = strings.ToLower(strings.TrimSpace(q))
		if q == "" {
			filtered = all
		} else {
			filtered = filtered[:0:0]
			for _, name := range all {
				if strings.Contains(name, q) {
					filtered = append(filtered, name)
				}
			}
		}
		status.SetText(fmt.Sprintf("%d icons", len(filtered)))
		grid.Refresh()
		grid.ScrollToTop()
	}

	w.SetContent(container.NewBorder(search, status, nil, nil, grid))
	w.Resize(fyne.NewSize(900, 600))
	w.ShowAndRun()
}
