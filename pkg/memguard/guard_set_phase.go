//ff:func feature=memguard type=command control=sequence
//ff:what Guard.SetPhase — name the phase the run is in; a new phase is sampled at once, before its work allocates
package memguard

import "time"

// SetPhase records phase for the log and StopError. A change takes a sample
// right away.
func (g *Guard) SetPhase(phase string) {
	if g == nil || *g.phase.Load() == phase {
		return
	}
	g.phase.Store(&phase)
	g.tick(time.Now())
}
