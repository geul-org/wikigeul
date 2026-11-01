//ff:func feature=memguard type=helper control=sequence
//ff:what TestMachineEase — the first call starts the count; before hold nothing; after it one step up; with every worker back only the level changes; a long quiet spell resets hold
package memguard

import "testing"

func TestMachineEase(t *testing.T) {
	m := fixtureMachine(4)
	m.level, m.active, m.hold, m.lastShrink = Caution, 3, 2*m.l.Hold, fixtureAt(0)
	var d decision
	m.ease(fixtureAt(1), 0, &d)
	if !m.calmSince.Equal(fixtureAt(1)) || m.level != Caution {
		t.Fatal("start of the count")
	}
	m.ease(fixtureAt(30), 0, &d)
	if m.level != Caution || m.active != 3 {
		t.Fatal("before hold")
	}
	m.ease(fixtureAt(61), 5, &d)
	if m.level != Normal || m.active != 4 || m.hold != 2*m.l.Hold || !m.lastRaise.Equal(fixtureAt(61)) {
		t.Fatalf("%+v", m)
	}
	if m.rssAt != 5 {
		t.Fatalf("rss at the raise %d", m.rssAt)
	}
	m.level = Caution
	m.ease(fixtureAt(700), 0, &d)
	if m.level != Normal || m.active != 4 || m.hold != m.l.Hold {
		t.Fatalf("%+v", m)
	}
}
