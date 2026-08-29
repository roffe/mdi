package all

import (
	"maps"
	"slices"
)

// Names returns the sorted names of all icons.
func Names() []string {
	return slices.Sorted(maps.Keys(Icons))
}
