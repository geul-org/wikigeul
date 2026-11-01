//ff:func feature=memguard type=helper control=sequence
//ff:what pressure.tighter — each slack the lower of the two; the description of the limit closer to its caution floor
package memguard

// tighter combines two limits.
func (p pressure) tighter(o pressure) pressure {
	d, r := min(p.Danger, o.Danger), min(p.Recover, o.Recover)
	if o.Caution < p.Caution {
		p = o
	}
	p.Danger, p.Recover = d, r
	return p
}
