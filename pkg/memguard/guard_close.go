//ff:func feature=memguard type=command control=sequence
//ff:what Guard.Close — stop sampling, put the Go heap limit back, wake waiting workers and log the peak
package memguard

import "fmt"

// Close ends the watch. It is safe on nil and may be deferred right after
// Start.
func (g *Guard) Close() {
	if g == nil {
		return
	}
	close(g.quit)
	<-g.done
	g.setMem(g.base)
	fmt.Fprintf(g.log, "memory peak rss=%dMiB\n", peakRSS(procStatus)/mib)
}
