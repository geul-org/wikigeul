//ff:func feature=memguard type=command control=sequence
//ff:what Guard.Ramp — work is about to start that takes memory per worker: go down to one worker and let them climb as memory answers
package memguard

// Ramp is called before workers start. Each worker's memory comes all at once
// on its first job (a compression encoder holds hundreds of MiB), and every
// worker starting together would pass a limit before any sample could see it.
// So they start one at a time, an interval apart, while the trend stays calm
// (machine.ease).
func (g *Guard) Ramp() {
	if g == nil {
		return
	}
	g.tickMu.Lock()
	defer g.tickMu.Unlock()
	if g.m.level == Danger {
		return
	}
	g.m.active, g.m.ramp, g.m.calmSince = 1, g.m.workers > 1, g.m.last
	g.setActive(1)
}
