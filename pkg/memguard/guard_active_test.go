//ff:func feature=memguard type=helper control=sequence
//ff:what TestGuardActive — nil gives want; otherwise the lower of want and the active workers, at least 1
package memguard

import "testing"

func TestGuardActive(t *testing.T) {
	var none *Guard
	if none.Active(7) != 7 {
		t.Fatal("nil")
	}
	g, _, _ := fixtureGuard(4)
	if g.Active(9) != 4 || g.Active(2) != 2 {
		t.Fatal(g.Active(9))
	}
	g.setActive(0)
	if g.Active(9) != 1 {
		t.Fatal(g.Active(9))
	}
}
