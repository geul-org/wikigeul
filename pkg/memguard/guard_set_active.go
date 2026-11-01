//ff:func feature=memguard type=command control=sequence
//ff:what Guard.setActive — change how many workers may run and tell whoever sizes itself by them; the waiting workers look again
package memguard

// setActive sets the count. OnActive hears counts of at least 1, after the
// gate changed, so its work (dropping encoders) meets fewer workers.
func (g *Guard) setActive(n int) {
	g.gate.L.Lock()
	old, hear := g.active, g.onCount
	g.active = n
	g.gate.L.Unlock()
	g.gate.Broadcast()
	if n != old && n > 0 && hear != nil {
		hear(n)
	}
}
