package all

import "testing"

func TestAll(t *testing.T) {
	if len(Icons) < 7000 {
		t.Fatalf("expected 7000+ icons, got %d", len(Icons))
	}
	for name, fn := range Icons {
		if len(fn().Content()) == 0 {
			t.Fatalf("icon %q has no content", name)
		}
	}
}
