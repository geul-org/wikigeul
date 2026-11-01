//ff:func feature=memguard type=helper control=sequence
//ff:what TestGuardOnActive — nil takes it and does nothing; a guard calls the last one registered; a count already down is told at once
package memguard

import "testing"

func TestGuardOnActive(t *testing.T) {
	var none *Guard
	none.OnActive(func(int) { t.Fatal("nil guard") })
	g, _, _ := fixtureGuard(4)
	got := 0
	g.OnActive(func(int) { t.Fatal("replaced") })
	g.OnActive(func(n int) { got = n })
	g.setActive(2)
	if got != 2 {
		t.Fatal(got)
	}
	late := 0
	g.OnActive(func(n int) { late = n })
	if late != 2 {
		t.Fatal(late)
	}
	g.setActive(0)
	g.OnActive(func(int) { t.Fatal("no worker to size for") })
}
