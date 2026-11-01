//ff:func feature=memguard type=helper control=sequence
//ff:what TestMachineStepOscillation — memory that answers the worker count: settles inside the band with one step down; where the band cannot hold a worker the steps down stay few and slow down
package memguard

import "testing"

func TestMachineStepOscillation(t *testing.T) {
	presses := func(base, each int64) (n, active int) {
		m := fixtureMachine(8)
		for s := 0; s < 3600; s++ {
			before := m.active
			d := m.step(fixtureAt(s), fixtureSample(100-base-each*int64(m.active)))
			if d.Active < before {
				n++
			}
		}
		return n, m.active
	}
	if n, active := presses(20, 8); n != 1 || active != 6 {
		t.Fatalf("band holds: %d steps down, ends at %d workers", n, active)
	}
	if n, _ := presses(10, 10); n > 8 {
		t.Fatalf("band too narrow for a worker: %d steps down in an hour", n)
	}
}
