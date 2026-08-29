package all

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestAll(t *testing.T) {
	test.NewApp()
	if len(Icons) < 7000 {
		t.Fatalf("expected 7000+ icons, got %d", len(Icons))
	}
	for name, fn := range Icons {
		if len(fn().Content()) == 0 {
			t.Fatalf("icon %q has no content", name)
		}
	}
}
