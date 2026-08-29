// Package mdi provides the Material Design Icons SVG set (v7.4.47,
// https://pictogrammers.com/library/mdi/) as Fyne resources.
//
// Each icon is its own function backed by its own //go:embed variable
// (generated in icons.go), so the linker only includes the icons you
// actually call — e.g. mdi.IconContentSave(). Wrap the result in
// theme.NewThemedResource to have it follow the Fyne theme's foreground
// color. Import github.com/roffe/mdi/all when you need lookup by name;
// it pins every icon into the binary.
package mdi

import "fyne.io/fyne/v2"

//go:generate go run gen.go

func res(name string, content []byte) fyne.Resource {
	return &fyne.StaticResource{StaticName: name, StaticContent: content}
}
