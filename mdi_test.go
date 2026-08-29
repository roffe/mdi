package mdi

import "testing"

func TestIcon(t *testing.T) {
	r := IconAccount()
	if r.Name() != "account.svg" {
		t.Fatalf("expected account.svg, got %s", r.Name())
	}
	if len(r.Content()) == 0 {
		t.Fatal("account icon has no content")
	}
}
