//ff:func feature=memguard type=helper control=sequence
//ff:what Guard.Phase — the phase the run is in
package memguard

// Phase is the name last given to SetPhase ("off" for no guard).
func (g *Guard) Phase() string {
	if g == nil {
		return "off"
	}
	return *g.phase.Load()
}
