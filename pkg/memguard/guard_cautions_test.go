//ff:func feature=memguard type=helper control=sequence
//ff:what TestGuardCautions — 0 for no guard, one for each entry into caution
package memguard

import "testing"

func TestGuardCautions(t *testing.T) {
	var none *Guard
	if none.Cautions() != 0 {
		t.Fatal("nil")
	}
	g, mem, _ := fixtureGuard(2)
	mem.s = fixtureSample(20)
	g.tick(fixtureAt(0))
	g.tick(fixtureAt(1))
	if g.Cautions() != 1 {
		t.Fatal(g.Cautions())
	}
}
