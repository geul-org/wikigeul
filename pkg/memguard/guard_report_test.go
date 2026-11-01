//ff:func feature=memguard type=helper control=iteration dimension=1
//ff:what TestGuardReport — the first line carries the state; quiet samples print nothing until the quiet stretch ends; a change of level, workers or phase prints
package memguard

import (
	"strings"
	"testing"
)

func TestGuardReport(t *testing.T) {
	g, mem, log := fixtureGuard(4)
	g.tick(fixtureAt(0))
	for _, want := range []string{"memory ", "phase=start", "level=normal", "avail=", "free=60.0%", "tight=system", "workers=4/4"} {
		if !strings.Contains(log.String(), want) {
			t.Fatalf("%q lacks %q", log, want)
		}
	}
	lines := func() int { return strings.Count(log.String(), "\n") }
	g.tick(fixtureAt(30))
	if lines() != 1 {
		t.Fatalf("quiet: %s", log)
	}
	g.tick(fixtureAt(61))
	if lines() != 2 {
		t.Fatalf("after the quiet stretch: %s", log)
	}
	g.SetPhase("convert")
	if lines() != 3 || !strings.Contains(log.String(), "phase=convert") {
		t.Fatalf("phase: %s", log)
	}
	mem.s = fixtureSample(20)
	g.tick(fixtureAt(70))
	if lines() != 4 || !strings.Contains(log.String(), "level=caution") || !strings.Contains(log.String(), "workers=2/4") {
		t.Fatalf("level: %s", log)
	}
}
