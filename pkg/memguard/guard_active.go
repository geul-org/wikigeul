//ff:func feature=memguard type=helper control=sequence
//ff:what Guard.Active — how many of want threads may run now, at least 1
package memguard

// Active bounds a thread count by the active workers.
func (g *Guard) Active(want int) int {
	if g == nil {
		return want
	}
	g.gate.L.Lock()
	defer g.gate.L.Unlock()
	return max(1, min(want, g.active))
}
