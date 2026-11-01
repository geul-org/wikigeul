//ff:func feature=memguard type=helper control=sequence
//ff:what TestGuardPhase — off for nil, the phase set otherwise
package memguard

import "testing"

func TestGuardPhase(t *testing.T) {
	var none *Guard
	if none.Phase() != "off" {
		t.Fatal(none.Phase())
	}
	g, _, _ := fixtureGuard(1)
	g.SetPhase("convert")
	if g.Phase() != "convert" {
		t.Fatal(g.Phase())
	}
}
