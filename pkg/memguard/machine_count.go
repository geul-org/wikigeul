//ff:func feature=memguard type=helper control=sequence
//ff:what machine.count — count the samples in a row at or heading across each floor: danger looks two intervals ahead, caution Lead ahead
package memguard

// count updates the confirmation counts. Stopping a run is expensive, so
// danger looks only as far ahead as the guard takes to act; slowing down is
// cheap, so caution looks Lead ahead.
func (m *machine) count(p pressure, fall float64) {
	m.dangerN = bump(m.dangerN, float64(p.Danger)-fall*2*m.l.Interval.Seconds() <= 0)
	m.cautionN = bump(m.cautionN, float64(p.Caution)-fall*m.l.Lead.Seconds() <= 0)
}
