//ff:func feature=memguard type=helper control=sequence
//ff:what TestGuardShare — want for no guard or plenty; a part of the budget when short; never below want/16
package memguard

import "testing"

func TestGuardShare(t *testing.T) {
	var none *Guard
	if none.Share(100, 4) != 100 {
		t.Fatal("nil")
	}
	g, _, _ := fixtureGuard(1)
	if g.Share(160, 4) != 160 {
		t.Fatal("plenty")
	}
	g.budget.Store(400)
	if got := g.Share(160, 4); got != 100 {
		t.Fatal(got)
	}
	g.budget.Store(-1)
	if got := g.Share(160, 4); got != 10 {
		t.Fatal(got)
	}
}
