//ff:func feature=memguard type=helper control=sequence
//ff:what TestMachineUpWait — Hold, an interval during a ramp
package memguard

import "testing"

func TestMachineUpWait(t *testing.T) {
	m := fixtureMachine(2)
	if m.upWait() != m.l.Hold {
		t.Fatal(m.upWait())
	}
	m.ramp = true
	if m.upWait() != m.l.Interval {
		t.Fatal(m.upWait())
	}
}
