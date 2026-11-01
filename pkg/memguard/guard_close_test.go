//ff:func feature=memguard type=command control=sequence
//ff:what TestGuardClose — nil does nothing; sampling ends, the Go heap limit is back and the peak is logged
package memguard

import (
	"strings"
	"testing"
)

func TestGuardClose(t *testing.T) {
	var none *Guard
	none.Close()
	g, mem, log := fixtureGuard(1)
	go g.run()
	g.Close()
	if mem.limits[len(mem.limits)-1] != g.base || !strings.Contains(log.String(), "memory peak rss=") {
		t.Fatalf("%v %s", mem.limits, log)
	}
}
