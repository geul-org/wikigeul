//ff:type feature=memguard type=model
//ff:what Guard — a running memory watch: the state machine and what the run asks it; a nil Guard is no guard, every method safe
package memguard

import (
	"io"
	"sync"
	"sync/atomic"
	"time"
)

// Guard watches memory while a run works (Start) and acts on what it finds:
// it limits the Go heap, gates the workers (Enter, Leave), gives memory back
// and, in danger, tells the run to stop (Stopped, Err). One goroutine samples;
// the run also samples at each phase change. tickMu keeps one tick at a time
// over m and the fields below it. gate guards active and inUse. A nil *Guard
// (Off) lets everything through.
type Guard struct {
	l       Limits
	read    Reader
	log     io.Writer
	workers int

	setMem  func(int64) int64
	reclaim func()

	tickMu     sync.Mutex
	m          machine
	base       int64
	memLimit   int64
	lastLog    time.Time
	lastLevel  Level
	lastActive int
	lastPhase  string
	reason     string

	gate    *sync.Cond
	active  int
	inUse   int
	onCount func(int)

	phase    atomic.Pointer[string]
	level    atomic.Int32
	spare    atomic.Int64
	budget   atomic.Int64
	near     atomic.Bool
	cautions atomic.Int64
	stop     atomic.Pointer[StopError]
	stopped  chan struct{}
	quit     chan struct{}
	done     chan struct{}
}
