//ff:func feature=memguard type=helper control=sequence
//ff:what Guard.OnActive — register who sizes itself by the active workers (the compressor's encoders), telling it the count at once when it is already down
package memguard

// OnActive makes f hear every new worker count. A run that starts under
// pressure has fewer active workers than it has, so f hears that count now.
// Later calls replace it.
func (g *Guard) OnActive(f func(int)) {
	if g == nil {
		return
	}
	g.gate.L.Lock()
	g.onCount = f
	n := g.active
	g.gate.L.Unlock()
	if f != nil && n > 0 && n < g.workers {
		f(n)
	}
}
