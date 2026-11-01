//ff:func feature=memguard type=command control=sequence
//ff:what Guard.tick — take one sample and act on the machine's decision; a failed reading is logged and changes nothing
package memguard

import (
	"fmt"
	"time"
)

// tick runs one step at now. Samples that cannot be read leave the state as
// it was: a guard that lost its numbers must not stop a healthy run.
func (g *Guard) tick(now time.Time) {
	g.tickMu.Lock()
	defer g.tickMu.Unlock()
	s, err := g.read(g.Phase())
	if err != nil {
		fmt.Fprintf(g.log, "memory %s read failed: %v\n", now.Format(time.RFC3339), err)
		return
	}
	d := g.m.step(now, s)
	g.apply(d)
	g.report(now, s, d)
}
