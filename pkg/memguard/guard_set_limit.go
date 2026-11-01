//ff:func feature=memguard type=command control=sequence
//ff:what Guard.setLimit — set the Go heap limit to the budget, above the live heap and under the ceiling; small changes are left out
package memguard

// setLimit keeps the Go runtime inside the budget. Under the live heap a
// limit would only make the collector run without end, so the limit stays
// an eighth above it (Danger and Caution deal with that case). The call
// stops the world, so a change of under 1/16 is not made.
func (g *Guard) setLimit(budget int64) {
	live := liveHeap()
	want := min(g.base, max(budget, live+live/8))
	if diff := want - g.memLimit; diff < g.memLimit/16 && -diff < g.memLimit/16 {
		return
	}
	g.memLimit = want
	g.setMem(want)
}
