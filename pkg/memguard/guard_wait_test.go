//ff:func feature=memguard type=helper control=sequence
//ff:what TestGuardWait — the interval, a quarter of it below the recovery floor
package memguard

import (
	"testing"
	"time"
)

func TestGuardWait(t *testing.T) {
	g, mem, _ := fixtureGuard(1)
	if g.wait() != time.Second {
		t.Fatal(g.wait())
	}
	mem.s = fixtureSample(33)
	g.tick(fixtureAt(0))
	if g.wait() != time.Second/4 {
		t.Fatal(g.wait())
	}
}
