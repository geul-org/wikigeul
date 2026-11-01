//ff:func feature=memguard type=helper control=sequence
//ff:what TestGuardLevel — nil is Normal; the level set by a step shows
package memguard

import "testing"

func TestGuardLevel(t *testing.T) {
	var none *Guard
	if none.Level() != Normal {
		t.Fatal("nil")
	}
	g, mem, _ := fixtureGuard(4)
	mem.s = fixtureSample(20)
	g.tick(fixtureAt(0))
	if g.Level() != Caution {
		t.Fatal(g.Level())
	}
}
