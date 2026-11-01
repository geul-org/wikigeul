//ff:func feature=memguard type=command control=sequence
//ff:what Guard.Pace — a safe point: note the phase and return the stop when there is one; what pack.StreamOptions.Pace wants
package memguard

// Pace is SetPhase and Err.
func (g *Guard) Pace(phase string) error {
	g.SetPhase(phase)
	return g.Err()
}
