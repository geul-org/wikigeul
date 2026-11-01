//ff:func feature=memguard type=command control=iteration dimension=1
//ff:what Guard.run — sample until Close: every Interval, a quarter of that while memory is short
package memguard

import "time"

// run is the sampling goroutine.
func (g *Guard) run() {
	defer close(g.done)
	for {
		select {
		case <-g.quit:
			return
		case <-time.After(g.wait()):
			g.tick(time.Now())
		}
	}
}
