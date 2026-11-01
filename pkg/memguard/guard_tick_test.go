//ff:func feature=memguard type=command control=sequence
//ff:what TestGuardTick — a sample steps the machine under the phase; a failed reading is logged and changes nothing
package memguard

import (
	"errors"
	"strings"
	"testing"
)

func TestGuardTick(t *testing.T) {
	g, mem, log := fixtureGuard(4)
	mem.s = fixtureSample(20)
	g.tick(fixtureAt(0))
	if g.Level() != Caution || mem.phases[0] != "start" {
		t.Fatalf("%v %v", g.Level(), mem.phases)
	}
	mem.err = errors.New("gone")
	mem.s = fixtureSample(1)
	g.tick(fixtureAt(1))
	if g.Level() != Caution || !strings.Contains(log.String(), "read failed: gone") {
		t.Fatalf("%v %s", g.Level(), log)
	}
}
