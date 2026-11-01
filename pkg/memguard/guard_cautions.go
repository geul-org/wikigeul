//ff:func feature=memguard type=helper control=sequence
//ff:what Guard.Cautions — how many times the run has entered Caution
package memguard

// Cautions lets a writer checkpoint once on each entry.
func (g *Guard) Cautions() int64 {
	if g == nil {
		return 0
	}
	return g.cautions.Load()
}
