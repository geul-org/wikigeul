//ff:func feature=memguard type=helper control=sequence
//ff:what TestMachineStepTrend — memory falling fast acts before the floor is crossed: Caution on the projection after two samples; one burst then flat brings no Danger
package memguard

import "testing"

func TestMachineStepTrend(t *testing.T) {
	m := fixtureMachine(8)
	m.step(fixtureAt(0), fixtureSample(80))
	if d := m.step(fixtureAt(1), fixtureSample(70)); d.Level != Normal {
		t.Fatalf("one sample of projection: %+v", d)
	}
	d := m.step(fixtureAt(2), fixtureSample(60))
	if d.Level != Caution || d.Active != 4 || d.Fall <= 0 {
		t.Fatalf("60%% free and falling: %+v", d)
	}
	b := fixtureMachine(8)
	b.step(fixtureAt(0), fixtureSample(80))
	b.step(fixtureAt(1), fixtureSample(20))
	if d = b.step(fixtureAt(2), fixtureSample(20)); d.Level == Danger {
		t.Fatalf("burst then flat: %+v", d)
	}
	f := fixtureMachine(8)
	f.step(fixtureAt(0), fixtureSample(60))
	f.step(fixtureAt(1), fixtureSample(45))
	f.step(fixtureAt(2), fixtureSample(30))
	if d = f.step(fixtureAt(3), fixtureSample(20)); d.Level != Danger {
		t.Fatalf("20%% free and falling fast, not yet at 15: %+v", d)
	}
}
