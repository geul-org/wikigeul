//ff:func feature=memguard type=helper control=sequence
//ff:what TestMachineCalm — calm needs the recovery floor now and in Lead at the present fall
package memguard

import "testing"

func TestMachineCalm(t *testing.T) {
	m := fixtureMachine(1)
	if m.calm(pressure{Recover: -1}, 0) {
		t.Fatal("below the floor")
	}
	if !m.calm(pressure{Recover: 100}, 1) {
		t.Fatal("above and slow")
	}
	if m.calm(pressure{Recover: 100}, 100) {
		t.Fatal("falling fast")
	}
}
