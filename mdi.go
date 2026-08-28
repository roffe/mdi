// Package mdi embeds the Material Design Icons SVG set (v7.4.47,
// https://pictogrammers.com/library/mdi/) as Fyne resources.
//
// Icons are addressed by their MDI name, e.g. "account", "chevron-down".
// Wrap the result in theme.NewThemedResource to have it follow the
// Fyne theme's foreground color.
package mdi

import (
	"embed"
	"sort"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//go:generate go run gen.go

//go:embed svg/*.svg
var svgs embed.FS

// IconName is an MDI icon name, e.g. "account". Constants for every
// icon are generated in names.go.
type IconName string

// Icon returns the named MDI icon as a Fyne resource. Unknown names
// fall back to the "border-none" placeholder.
func Icon(name IconName) fyne.Resource {
	b, err := svgs.ReadFile("svg/" + string(name) + ".svg")
	if err != nil {
		return Icon(IconBorderNone)
	}
	return &fyne.StaticResource{StaticName: string(name) + ".svg", StaticContent: b}
}

// ThemedIcon returns the named MDI icon colored by the current theme's
// foreground. Unknown names fall back to the "border-none" placeholder.
func ThemedIcon(name IconName) fyne.Resource {
	return theme.NewThemedResource(Icon(name))
}

// Names returns the sorted names of all embedded icons.
func Names() []string {
	entries, _ := svgs.ReadDir("svg")
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		names = append(names, strings.TrimSuffix(e.Name(), ".svg"))
	}
	sort.Strings(names)
	return names
}
