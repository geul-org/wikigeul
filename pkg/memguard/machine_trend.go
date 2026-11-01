//ff:func feature=memguard type=helper control=sequence
//ff:what machine.trend — how fast the slack is falling, bytes a second: the lower of the smoothed and the last rate, so a step that ended is no trend; samples closer than half an interval do not count
package memguard

import "time"

// trend updates the rates and returns the fall (0 when rising). The smoothed
// rate outlives a single step of memory, and a step that has ended must not
// read as a slope: the fall counts only while the last step fell too. A
// sample taken right after another (at a phase change) would turn a small
// step into a huge rate, so it leaves the rates as they are.
func (m *machine) trend(now time.Time, slack int64) float64 {
	if m.last.IsZero() {
		m.last, m.prev = now, slack
		return 0
	}
	if dt := now.Sub(m.last); dt >= m.l.Interval/2 {
		m.rate = float64(m.prev-slack) / dt.Seconds()
		m.fall = m.fall/2 + m.rate/2
		m.last, m.prev = now, slack
	}
	return max(0, min(m.fall, m.rate))
}
