package mdi

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestIcon(t *testing.T) {
	test.NewApp()
	r := IconAccount()
	if r.Name() != "foreground_account.svg" {
		t.Fatalf("expected foreground_account.svg, got %s", r.Name())
	}
	if len(r.Content()) == 0 {
		t.Fatal("account icon has no content")
	}
}
