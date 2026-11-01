//ff:func feature=memguard type=helper control=iteration dimension=1
//ff:what TestMachineStepRamp — workers climb one an interval while memory is calm, wait while it falls, stop climbing where the next one would not fit, and pressure ends the ramp
package memguard

import "testing"

func TestMachineStepRamp(t *testing.T) {
	m := fixtureMachine(4)
	m.active, m.ramp = 1, true
	m.step(fixtureAt(0), fixtureSample(80))
	for s := 1; s <= 6; s++ {
		m.step(fixtureAt(s), fixtureSample(80))
	}
	if m.active != 4 || m.ramp {
		t.Fatalf("%d workers, ramp %v", m.active, m.ramp)
	}
	c := fixtureMachine(4)
	c.active, c.ramp = 1, true
	rss := []int64{1, 1, 1, 7, 7, 7, 7, 7}
	s := Sample{Total: 64 << 30}
	for i, gib := range rss {
		s.RSS, s.Avail = gib<<30, (32-gib)<<30
		c.step(fixtureAt(10*i), s)
	}
	if c.active != 3 || c.cost != 6<<30 {
		t.Fatalf("a worker costs 6 GiB and under 3 GiB is above the floor: %d workers, cost %d", c.active, c.cost)
	}
	p := fixtureMachine(4)
	p.active, p.ramp = 1, true
	p.step(fixtureAt(0), fixtureSample(20))
	if p.ramp {
		t.Fatal("pressure ends the ramp")
	}
}
