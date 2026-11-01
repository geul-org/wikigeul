//ff:func feature=memguard type=command control=sequence
//ff:what Start — begin watching memory for a run of workers workers: first sample now, then a goroutine; nil when the guard is off
package memguard

import (
	"fmt"
	"io"
	"time"
)

// Start checks that memory can be read and starts sampling, logging to log.
// With l.Off it returns nil, a Guard that does nothing. Close ends it.
func Start(l Limits, workers int, log io.Writer) (*Guard, error) {
	if l.Off {
		return nil, nil
	}
	read := l.Read
	if read == nil {
		read = procReader(l.Meminfo)
	}
	if _, err := read("start"); err != nil {
		return nil, fmt.Errorf("memory guard: %w (-mem-off runs without it)", err)
	}
	g := newGuard(l, workers, read, log)
	g.tick(time.Now())
	go g.run()
	return g, nil
}
