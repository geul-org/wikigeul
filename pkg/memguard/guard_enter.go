//ff:func feature=memguard type=helper control=iteration dimension=1
//ff:what Guard.Enter — a worker asks to start one job: wait while the active workers are all busy; false when the run is stopping
package memguard

// Enter holds a job back until fewer than the active count are running.
// Every worker that entered must Leave. The wait is in arrival order, so the
// job that came first runs first. A false worker takes no more jobs.
func (g *Guard) Enter() bool {
	if g == nil {
		return true
	}
	g.gate.L.Lock()
	defer g.gate.L.Unlock()
	for g.stop.Load() == nil && g.inUse >= g.active {
		g.gate.Wait()
	}
	if g.stop.Load() != nil {
		return false
	}
	g.inUse++
	return true
}
