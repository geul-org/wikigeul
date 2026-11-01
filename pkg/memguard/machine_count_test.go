//ff:func feature=memguard type=helper control=sequence
//ff:what TestMachineCount — below a floor counts at once; falling, caution looks further ahead than danger; above both resets
package memguard

import "testing"

func TestMachineCount(t *testing.T) {
	m := fixtureMachine(1)
	m.count(pressure{Danger: -1, Caution: -1}, 0)
	if m.dangerN != 1 || m.cautionN != 1 {
		t.Fatalf("%d %d", m.dangerN, m.cautionN)
	}
	m.count(pressure{Danger: 100, Caution: 100}, 20)
	if m.dangerN != 0 || m.cautionN != 2 {
		t.Fatalf("20 a second: %d %d", m.dangerN, m.cautionN)
	}
	m.count(pressure{Danger: 100, Caution: 100}, 50)
	if m.dangerN != 1 {
		t.Fatal("50 a second reaches danger in two intervals")
	}
	m.count(pressure{Danger: 100, Caution: 100}, 0)
	if m.dangerN != 0 || m.cautionN != 0 {
		t.Fatalf("%d %d", m.dangerN, m.cautionN)
	}
}
