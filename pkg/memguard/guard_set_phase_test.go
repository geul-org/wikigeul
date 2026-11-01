//ff:func feature=memguard type=command control=sequence
//ff:what TestGuardSetPhase — nil does nothing; a new phase samples at once, the same phase does not
package memguard

import "testing"

func TestGuardSetPhase(t *testing.T) {
	var none *Guard
	none.SetPhase("x")
	g, mem, _ := fixtureGuard(1)
	g.SetPhase("sort:scan")
	g.SetPhase("sort:scan")
	if len(mem.phases) != 1 || mem.phases[0] != "sort:scan" {
		t.Fatalf("%v", mem.phases)
	}
}
