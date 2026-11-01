//ff:func feature=memguard type=helper control=sequence
//ff:what Guard.Stopped — a channel closed when the run must stop; nil, never ready, for no guard
package memguard

// Stopped lets a writer wait for a result and for the stop at once.
func (g *Guard) Stopped() <-chan struct{} {
	if g == nil {
		return nil
	}
	return g.stopped
}
