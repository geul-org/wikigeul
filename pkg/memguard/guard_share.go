//ff:func feature=memguard type=helper control=sequence
//ff:what Guard.Share — want bytes, no more than a part of the process's budget, no less than want/shareFloor
package memguard

// Share bounds a buffer that sets its own size (the sort's held records) by
// budget/part. The floor keeps a short budget from making a run of every
// line.
func (g *Guard) Share(want, part int64) int64 {
	if g == nil {
		return want
	}
	return max(want/shareFloor, min(want, g.budget.Load()/part))
}
