//ff:func feature=memguard type=helper control=sequence
//ff:what machine.calm — whether memory stands above the recovery floor and will still in Lead at the present trend
package memguard

// calm is true when stepping up is safe to consider.
func (m *machine) calm(p pressure, fall float64) bool {
	return p.Recover >= 0 && float64(p.Recover)-fall*m.l.Lead.Seconds() > 0
}
