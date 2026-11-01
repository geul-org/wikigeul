//ff:func feature=memguard type=helper control=sequence
//ff:what machine.upWait — the calm needed before a step up: an interval while the workers climb at the start, Hold after that
package memguard

import "time"

// upWait is short during a ramp: nothing has gone wrong yet, and a sample
// sees what the last worker added.
func (m *machine) upWait() time.Duration {
	if m.ramp {
		return m.l.Interval
	}
	return m.hold
}
