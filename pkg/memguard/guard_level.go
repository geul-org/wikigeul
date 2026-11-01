//ff:func feature=memguard type=helper control=sequence
//ff:what Guard.Level — the guard's level now
package memguard

// Level is Normal for no guard.
func (g *Guard) Level() Level {
	if g == nil {
		return Normal
	}
	return Level(g.level.Load())
}
