//ff:func feature=memguard type=helper control=sequence
//ff:what Guard.wait — the time to the next sample: shorter while memory stands below the recovery floor
package memguard

import "time"

// wait is Interval, or a quarter of it when memory is short, when what
// happens next matters most.
func (g *Guard) wait() time.Duration {
	if g.near.Load() {
		return g.l.Interval / 4
	}
	return g.l.Interval
}
