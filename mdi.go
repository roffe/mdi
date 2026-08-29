// Package mdi provides the Material Design Icons SVG set (v7.4.47,
// https://pictogrammers.com/library/mdi/) as Fyne resources.
//
// Each icon is its own function backed by its own //go:embed variable
// (generated in icons.go), so the linker only includes the icons you
// actually call — e.g. mdi.IconContentSave(). Icons are themed: they
// follow the Fyne theme's foreground color. Import github.com/roffe/mdi/all
// when you need lookup by name; it pins every icon into the binary.
package mdi

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//go:generate go run gen.go

func res(name string, content []byte) fyne.Resource {
	return theme.NewThemedResource(&fyne.StaticResource{StaticName: name, StaticContent: content})
}
