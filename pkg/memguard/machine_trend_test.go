//ff:func feature=memguard type=helper control=sequence
//ff:what TestMachineTrend — first sample no trend; a fall read in bytes a second; a sample too soon leaves it; a flat step ends it; a rise reads 0
package memguard

import (
	"testing"
	"time"
)

func TestMachineTrend(t *testing.T) {
	m := fixtureMachine(1)
	if f := m.trend(fixtureAt(0), 1000); f != 0 {
		t.Fatal(f)
	}
	if f := m.trend(fixtureAt(2), 800); f != 50 || m.rate != 100 {
		t.Fatalf("%v %v", f, m.rate)
	}
	if f := m.trend(fixtureAt(2).Add(time.Millisecond), 0); f != 50 || m.prev != 800 {
		t.Fatalf("too soon: %v", f)
	}
	if f := m.trend(fixtureAt(3), 800); f != 0 || m.fall != 25 {
		t.Fatalf("flat: %v %v", f, m.fall)
	}
	if f := m.trend(fixtureAt(4), 2000); f != 0 || m.rate >= 0 {
		t.Fatalf("rising: %v", f)
	}
}
