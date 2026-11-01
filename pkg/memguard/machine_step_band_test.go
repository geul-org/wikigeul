//ff:func feature=memguard type=helper control=iteration dimension=1
//ff:what TestMachineStepBand — between the caution and recovery floors nothing moves: no step up or down however long it lasts, and the calm count starts over
package memguard

import "testing"

func TestMachineStepBand(t *testing.T) {
	m := fixtureMachine(8)
	m.step(fixtureAt(0), fixtureSample(29))
	for s := 1; s < 200; s++ {
		d := m.step(fixtureAt(s), fixtureSample(33))
		if d.Level != Caution || d.Active != 4 || d.Reclaim {
			t.Fatalf("second %d: %+v", s, d)
		}
	}
	m.step(fixtureAt(200), fixtureSample(50))
	m.step(fixtureAt(201), fixtureSample(50))
	if m.calmSince.IsZero() {
		t.Fatal("calm not counted")
	}
	if m.step(fixtureAt(202), fixtureSample(33)); !m.calmSince.IsZero() {
		t.Fatal("calm count not started over")
	}
}
