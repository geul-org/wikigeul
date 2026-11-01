//ff:func feature=memguard type=helper control=iteration dimension=1
//ff:what TestMachinePress — a step down halves the workers but never below one; a step down soon after a raise stretches Hold up to its bound; one long after does not
package memguard

import "testing"

func TestMachinePress(t *testing.T) {
	m := fixtureMachine(2)
	var d decision
	m.press(fixtureAt(0), &d)
	if m.active != 1 || !d.Reclaim || m.level != Caution {
		t.Fatalf("%+v", m)
	}
	d = decision{}
	if m.press(fixtureAt(1), &d); d.Reclaim {
		t.Fatal("within Cooldown")
	}
	m.press(fixtureAt(10), &d)
	if m.active != 1 || !d.Reclaim {
		t.Fatal("reclaim goes on at one worker")
	}
	m.lastRaise = fixtureAt(19)
	m.press(fixtureAt(20), &d)
	if m.hold != 2*m.l.Hold {
		t.Fatalf("hold %v", m.hold)
	}
	for s := 30; s < 100; s += 10 {
		m.lastRaise = fixtureAt(s)
		m.press(fixtureAt(s), &d)
	}
	if m.hold != maxHoldFactor*m.l.Hold {
		t.Fatalf("hold %v", m.hold)
	}
	n := fixtureMachine(2)
	n.lastRaise = fixtureAt(0)
	n.press(fixtureAt(1000), &d)
	if n.hold != n.l.Hold {
		t.Fatalf("hold %v", n.hold)
	}
}
