//ff:func feature=memguard type=helper control=sequence
//ff:what Guard.Leave — a worker finished its job: the next waiting one may start
package memguard

// Leave ends what Enter began.
func (g *Guard) Leave() {
	if g == nil {
		return
	}
	g.gate.L.Lock()
	g.inUse--
	g.gate.L.Unlock()
	g.gate.Signal()
}
