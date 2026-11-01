//ff:func feature=memguard type=helper control=sequence
//ff:what Limits.check — the settings must make sense: shares between 0 and 1, danger below caution, a positive interval
package memguard

import "errors"

// check refuses limits a guard cannot keep.
func (l Limits) check() error {
	if l.MinFree <= 0 || l.MinFree+l.Gap >= 1 {
		return errors.New("memguard: -mem-min-free must be above 0 and leave room to recover below 100 (-mem-off turns the guard off)")
	}
	if l.DangerFree < 0 || l.DangerFree > l.MinFree {
		return errors.New("memguard: -mem-danger-free must be between 0 and -mem-min-free")
	}
	if l.Max < 0 || l.Interval <= 0 {
		return errors.New("memguard: -mem-max must not be negative and -mem-interval must be positive")
	}
	return nil
}
