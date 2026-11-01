//ff:func feature=memguard type=helper control=sequence
//ff:what newGuard — a Guard over a Reader, not started: every worker active, the Go heap limit as the environment set it as ceiling
package memguard

import (
	"io"
	"math"
	"runtime/debug"
	"sync"
)

// newGuard makes a guard for workers workers. The Go memory limit at this
// moment (GOMEMLIMIT, or none) stays the ceiling of what it sets.
func newGuard(l Limits, workers int, read Reader, log io.Writer) *Guard {
	l = l.withDefaults()
	g := &Guard{l: l, read: read, log: log, workers: max(1, workers), setMem: debug.SetMemoryLimit, reclaim: debug.FreeOSMemory,
		m: newMachine(l, workers), active: max(1, workers), lastActive: max(1, workers), gate: sync.NewCond(&sync.Mutex{}),
		stopped: make(chan struct{}), quit: make(chan struct{}), done: make(chan struct{})}
	g.base = g.setMem(-1)
	g.memLimit = g.base
	g.budget.Store(math.MaxInt64)
	g.spare.Store(math.MaxInt64)
	phase := "start"
	g.phase.Store(&phase)
	return g
}
