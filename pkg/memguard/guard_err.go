//ff:func feature=memguard type=helper control=sequence
//ff:what Guard.Err — the StopError once the guard has said stop, else nil
package memguard

// Err is what a safe point returns.
func (g *Guard) Err() error {
	if g == nil {
		return nil
	}
	if e := g.stop.Load(); e != nil {
		return *e
	}
	return nil
}
