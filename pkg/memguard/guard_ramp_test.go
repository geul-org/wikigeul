//ff:func feature=memguard type=command control=iteration dimension=1
//ff:what TestGuardRamp — nil does nothing; one worker from now and whoever sizes by them hears it; workers climb from samples; a stopped run stays stopped
package memguard

import "testing"

func TestGuardRamp(t *testing.T) {
	var none *Guard
	none.Ramp()
	g, mem, _ := fixtureGuard(3)
	heard := 0
	g.OnActive(func(n int) { heard = n })
	g.Ramp()
	if g.Active(3) != 1 || heard != 1 || !g.m.ramp {
		t.Fatalf("%d %d", g.Active(3), heard)
	}
	for s := 0; s < 8; s++ {
		g.tick(fixtureAt(s))
	}
	if heard != 3 {
		t.Fatal(heard)
	}
	mem.s = fixtureSample(5)
	g.tick(fixtureAt(9))
	g.tick(fixtureAt(10))
	g.Ramp()
	if g.Err() == nil || g.m.active != 0 {
		t.Fatal("stop undone")
	}
	single, _, _ := fixtureGuard(1)
	if single.Ramp(); single.m.ramp {
		t.Fatal("nothing to climb")
	}
}
