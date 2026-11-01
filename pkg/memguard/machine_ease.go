//ff:func feature=memguard type=helper control=sequence
//ff:what machine.ease — after the wait of calm, one step up: Caution back to Normal, one more worker if what a worker costs fits above the recovery floor; a long quiet spell brings Hold back to its setting
package memguard

import (
	"fmt"
	"time"
)

// ease counts the calm stretch and, after upWait, takes one step up and
// starts the count again, so workers come back one by one. A worker comes
// back only if the most a raise has cost still leaves memory above the
// recovery floor: a raise that cannot fit is not tried.
func (m *machine) ease(now time.Time, rss int64, d *decision) {
	if m.calmSince.IsZero() {
		m.calmSince = now
		return
	}
	if now.Sub(m.calmSince) < m.upWait() {
		return
	}
	m.calmSince = now
	if now.Sub(m.lastShrink) > 10*m.hold {
		m.hold = m.l.Hold
	}
	if m.level == Caution {
		m.level = Normal
		d.Reason = fmt.Sprintf("recovered: %s free %s", d.Press.Name, freeText(d.Press))
	}
	if m.active < m.workers && d.Press.Recover >= m.cost {
		m.active++
		m.lastRaise, m.rssAt = now, rss
		m.ramp = m.ramp && m.active < m.workers
		d.Reason = fmt.Sprintf("raising workers: %s free %s", d.Press.Name, freeText(d.Press))
	}
}
