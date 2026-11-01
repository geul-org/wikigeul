//ff:func feature=memguard type=helper control=sequence
//ff:what machine.stop — enter Danger: no worker runs and the run stops at its next safe point
package memguard

import "fmt"

// stop latches Danger: the safe point is near, and a run that stops is
// resumed by whoever runs it, so a brief recovery must not undo it.
func (m *machine) stop(d *decision) {
	m.level, m.active = Danger, 0
	d.Reclaim = true
	d.Reason = fmt.Sprintf("%s free %s at the danger floor", d.Press.Name, freeText(d.Press))
}
