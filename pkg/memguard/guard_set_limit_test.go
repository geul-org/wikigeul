//ff:func feature=memguard type=command control=sequence
//ff:what TestGuardSetLimit — the budget becomes the limit under the ceiling; small changes are skipped; never under the live heap
package memguard

import "testing"

func TestGuardSetLimit(t *testing.T) {
	g, mem, _ := fixtureGuard(1)
	g.setLimit(1 << 40)
	if len(mem.limits) != 1 || mem.limits[0] != 1<<40 {
		t.Fatalf("%v", mem.limits)
	}
	g.setLimit(1<<40 + 1<<30)
	g.setLimit(1<<40 - 1<<30)
	if len(mem.limits) != 1 {
		t.Fatalf("small changes: %v", mem.limits)
	}
	g.setLimit(0)
	if got := mem.limits[len(mem.limits)-1]; got < liveHeap() {
		t.Fatalf("limit %d under the live heap", got)
	}
	g.base = 1 << 35
	g.setLimit(1 << 40)
	if got := mem.limits[len(mem.limits)-1]; got != 1<<35 {
		t.Fatalf("ceiling: %d", got)
	}
}
