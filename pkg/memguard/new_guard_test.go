//ff:func feature=memguard type=helper control=sequence
//ff:what TestNewGuard — every worker active, at least one; the Go memory limit in force is the ceiling; plenty until the first sample
package memguard

import (
	"io"
	"math"
	"runtime/debug"
	"testing"
)

func TestNewGuard(t *testing.T) {
	g := newGuard(Limits{}, 0, nil, io.Discard)
	if g.active != 1 || g.workers != 1 || g.base != debug.SetMemoryLimit(-1) || g.Phase() != "start" || g.budget.Load() != math.MaxInt64 {
		t.Fatalf("%+v", g)
	}
}
