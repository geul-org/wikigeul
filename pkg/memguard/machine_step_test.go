//ff:func feature=memguard type=helper control=sequence
//ff:what TestMachineStep — Normal above the floor; Caution below it halves the workers and reclaims; Danger needs two samples below its floor, takes every worker and stays
package memguard

import "testing"

func TestMachineStep(t *testing.T) {
	m := fixtureMachine(8)
	if d := m.step(fixtureAt(0), fixtureSample(60)); d.Level != Normal || d.Active != 8 || d.Reclaim || d.Near {
		t.Fatalf("%+v", d)
	}
	m.step(fixtureAt(1), fixtureSample(60))
	d := m.step(fixtureAt(2), fixtureSample(29))
	if d.Level != Caution || d.Active != 4 || !d.Reclaim || !d.Near || d.Reason == "" || d.Press.Name != "system" {
		t.Fatalf("%+v", d)
	}
	if d = m.step(fixtureAt(3), fixtureSample(29)); d.Level != Caution || d.Reclaim || d.Active != 4 {
		t.Fatalf("within Cooldown: %+v", d)
	}
	if d = m.step(fixtureAt(4), fixtureSample(14)); d.Level != Caution {
		t.Fatalf("one sample below danger: %+v", d)
	}
	if d = m.step(fixtureAt(5), fixtureSample(14)); d.Level != Danger || d.Active != 0 || d.Reason == "" {
		t.Fatalf("%+v", d)
	}
	if d = m.step(fixtureAt(100), fixtureSample(90)); d.Level != Danger || d.Active != 0 {
		t.Fatalf("danger must stay: %+v", d)
	}
}
