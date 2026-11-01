//ff:func feature=memguard type=command control=sequence
//ff:what TestGuardRun — the goroutine samples until Close, reaching a stop by itself when memory runs out
package memguard

import (
	"io"
	"sync/atomic"
	"testing"
	"time"
)

func TestGuardRun(t *testing.T) {
	var low atomic.Bool
	read := func(string) (Sample, error) {
		if low.Load() {
			return fixtureSample(10), nil
		}
		return fixtureSample(60), nil
	}
	g := newGuard(Limits{Interval: time.Millisecond}, 4, read, io.Discard)
	go g.run()
	defer g.Close()
	low.Store(true)
	select {
	case <-g.Stopped():
	case <-time.After(5 * time.Second):
		t.Fatal("no stop")
	}
}
