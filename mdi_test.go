package mdi

import "testing"

func TestIcon(t *testing.T) {
	if Icon("account") == nil {
		t.Fatal("account icon missing")
	}
	if Icon(IconAccount) == nil {
		t.Fatal("IconAccount constant broken")
	}
	if got := Icon("no-such-icon-xyz").Name(); got != "border-none.svg" {
		t.Fatalf("expected border-none fallback, got %s", got)
	}
	if n := len(Names()); n < 7000 {
		t.Fatalf("expected 7000+ icons, got %d", n)
	}
}
