//ff:func feature=memguard type=helper control=selection
//ff:what machine.step — one sample: danger stays; a confirmed crossing of the danger floor stops; below the caution floor or heading there presses; a calm stretch eases; between the floors nothing moves
package memguard

import "time"

// step decides from a sample at now. Stepping down is cheap to undo and
// happens at once below the floor; stepping up needs memory above the
// higher recovery floor for a while, and the band between the two floors
// changes nothing, so the guard does not swing around one line.
func (m *machine) step(now time.Time, s Sample) decision {
	p := pressureOf(m.l, s)
	fall := m.trend(now, p.Caution)
	m.count(p, fall)
	if m.rssAt > 0 {
		m.cost = max(m.cost, s.RSS-m.rssAt)
	}
	d := decision{Press: p, Fall: fall, Budget: s.RSS + p.Caution, Near: p.Recover < 0}
	switch {
	case m.level == Danger:
	case m.dangerN >= m.l.Confirm:
		m.stop(&d)
	case p.Caution <= 0 || m.cautionN >= m.l.Confirm:
		m.press(now, &d)
	case m.calm(p, fall):
		m.ease(now, s.RSS, &d)
	default:
		m.calmSince = time.Time{}
	}
	d.Level, d.Active = m.level, m.active
	return d
}
