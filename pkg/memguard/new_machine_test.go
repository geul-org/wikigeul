//ff:func feature=memguard type=helper control=sequence
//ff:what TestNewMachine — every worker active, at least one, Hold at its setting
package memguard

import "testing"

func TestNewMachine(t *testing.T) {
	m := newMachine(Limits{}.withDefaults(), 0)
	if m.workers != 1 || m.active != 1 || m.hold != m.l.Hold || m.level != Normal {
		t.Fatalf("%+v", m)
	}
}
