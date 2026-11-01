//ff:func feature=memguard type=helper control=iteration dimension=1
//ff:what TestMachineStepRecover — after Hold above the recovery floor the level returns to Normal and workers come back one at a time, each after another Hold
package memguard

import "testing"

func TestMachineStepRecover(t *testing.T) {
	m := fixtureMachine(8)
	m.step(fixtureAt(0), fixtureSample(29))
	var rises []int
	for s := 1; s <= 200; s++ {
		before := m.active
		d := m.step(fixtureAt(s), fixtureSample(50))
		if d.Active > before+1 {
			t.Fatalf("second %d: %d to %d", s, before, d.Active)
		}
		if d.Active > before {
			rises = append(rises, s)
		}
		if s < 30 && d.Level == Normal {
			t.Fatalf("second %d: recovered before Hold", s)
		}
	}
	if m.level != Normal || m.active != 8 || len(rises) != 4 || rises[1]-rises[0] != 30 {
		t.Fatalf("level %v active %d rises %v", m.level, m.active, rises)
	}
}
