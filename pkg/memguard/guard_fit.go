//ff:func feature=memguard type=helper control=sequence
//ff:what Guard.Fit — how many buffers of size bytes fit in half the memory above the caution floor, between lo and hi
package memguard

// Fit sizes a merge to the memory there is now: hi with no guard or plenty,
// lo with none to spare.
func (g *Guard) Fit(size int64, lo, hi int) int {
	if g == nil {
		return hi
	}
	return int(min(max(g.spare.Load()/2/size, int64(lo)), int64(hi)))
}
