//ff:func feature=memguard type=helper control=sequence
//ff:what Guard.report — print a state line when level, workers or phase changed, or after a quiet stretch
package memguard

import (
	"fmt"
	"time"
)

// report prints "memory <time> phase=… level=… rss=… avail=… free=… tight=…
// go-limit=… workers=a/n fall=… reason=…", a line a watcher can parse.
func (g *Guard) report(now time.Time, s Sample, d decision) {
	phase := g.Phase()
	changed := d.Level != g.lastLevel || d.Active != g.lastActive || phase != g.lastPhase || g.lastLog.IsZero()
	if !changed && now.Sub(g.lastLog) < logEverySamples*g.l.Interval {
		return
	}
	g.lastLevel, g.lastActive, g.lastPhase, g.lastLog = d.Level, d.Active, phase, now
	fmt.Fprintf(g.log, "memory %s phase=%s level=%s rss=%dMiB avail=%dMiB/%dMiB free=%s tight=%s go-limit=%dMiB workers=%d/%d fall=%dMiB/s reason=%q\n",
		now.Format(time.RFC3339), phase, d.Level, s.RSS/mib, s.Avail/mib, s.Total/mib, freeText(d.Press), d.Press.Name, g.memLimit/mib, d.Active, g.workers, int64(d.Fall)/mib, g.reason)
}
