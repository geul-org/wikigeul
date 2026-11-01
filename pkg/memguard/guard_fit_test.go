//ff:func feature=memguard type=helper control=sequence
//ff:what TestGuardFit — hi for no guard; half the memory above the caution floor in units, between lo and hi; lo below the floor
package memguard

import "testing"

func TestGuardFit(t *testing.T) {
	var none *Guard
	if none.Fit(1, 2, 64) != 64 {
		t.Fatal("nil")
	}
	g, _, _ := fixtureGuard(1)
	if got := g.Fit(1<<20, 8, 64); got != 64 {
		t.Fatal("before any sample:", got)
	}
	g.spare.Store(40 << 20)
	if got := g.Fit(1<<20, 8, 64); got != 20 {
		t.Fatal(got)
	}
	g.spare.Store(-1)
	if got := g.Fit(1<<20, 8, 64); got != 8 {
		t.Fatal(got)
	}
}
