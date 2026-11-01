//ff:func feature=memguard type=helper control=sequence
//ff:what machine.press — Caution: halve the active workers and give memory back, at most once a Cooldown; a raise that went wrong stretches Hold
package memguard

import (
	"fmt"
	"time"
)

// press is the step down. The level changes at once; the workers and the
// reclaim wait for Cooldown since the last step, since memory takes time to
// answer. Halving (with one step up at a time) settles quickly; a step down
// soon after a raise means that raise was too many, so Hold doubles, up to
// maxHoldFactor times, and the next raises come later.
func (m *machine) press(now time.Time, d *decision) {
	if m.level != Caution {
		d.Reason = fmt.Sprintf("%s free %s, losing %d MiB/s", d.Press.Name, freeText(d.Press), int64(d.Fall)/mib)
	}
	m.level, m.ramp = Caution, false
	m.calmSince = time.Time{}
	if !m.lastShrink.IsZero() && now.Sub(m.lastShrink) < m.l.Cooldown {
		return
	}
	m.lastShrink, d.Reclaim = now, true
	m.active = max(1, m.active/2)
	if !m.lastRaise.IsZero() && now.Sub(m.lastRaise) < 2*m.hold {
		m.hold = min(2*m.hold, maxHoldFactor*m.l.Hold)
	}
}
