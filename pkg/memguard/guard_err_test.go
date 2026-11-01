//ff:func feature=memguard type=helper control=sequence
//ff:what TestGuardErr — nil for no guard and before a stop, the StopError after
package memguard

import "testing"

func TestGuardErr(t *testing.T) {
	var none *Guard
	if none.Err() != nil {
		t.Fatal("nil")
	}
	g, _, _ := fixtureGuard(1)
	if g.Err() != nil {
		t.Fatal("before")
	}
	g.apply(decision{Level: Danger, Press: pressure{Size: 1}})
	if _, ok := g.Err().(StopError); !ok {
		t.Fatal(g.Err())
	}
}
