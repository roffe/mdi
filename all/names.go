package all

import "sort"

// Names returns the sorted names of all icons.
func Names() []string {
	names := make([]string, 0, len(Icons))
	for name := range Icons {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
