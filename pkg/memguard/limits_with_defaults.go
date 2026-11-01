//ff:func feature=memguard type=helper control=sequence
//ff:what Limits.withDefaults — the limits with every zero setting at its default, times following the interval
package memguard

import "time"

// withDefaults fills what was left zero (Limits).
func (l Limits) withDefaults() Limits {
	if l.MinFree == 0 {
		l.MinFree = DefaultMinFree
	}
	if l.DangerFree == 0 {
		l.DangerFree = l.MinFree / 2
	}
	if l.Gap == 0 {
		l.Gap = defaultGap
	}
	if l.Interval == 0 {
		l.Interval = time.Second
	}
	if l.Lead == 0 {
		l.Lead = 10 * l.Interval
	}
	if l.Hold == 0 {
		l.Hold = 30 * l.Interval
	}
	if l.Cooldown == 0 {
		l.Cooldown = 5 * l.Interval
	}
	if l.Confirm == 0 {
		l.Confirm = defaultConfirm
	}
	return l
}
