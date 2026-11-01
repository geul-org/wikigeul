//ff:func feature=memguard type=helper control=sequence
//ff:what fixtureGuard — a guard not started, over a fixtureMem, its log in a buffer, its heap limit calls recorded
package memguard

import (
	"bytes"
	"math"
)

// fixtureGuard makes a guard at the default limits with no ceiling.
func fixtureGuard(workers int) (*Guard, *fixtureMem, *bytes.Buffer) {
	mem := &fixtureMem{s: fixtureSample(60)}
	log := &bytes.Buffer{}
	read := func(phase string) (Sample, error) {
		mem.phases = append(mem.phases, phase)
		return mem.s, mem.err
	}
	g := newGuard(Limits{}, workers, read, log)
	g.base, g.memLimit = math.MaxInt64, math.MaxInt64
	g.reclaim = func() { mem.reclaims++ }
	g.setMem = func(n int64) int64 { mem.limits = append(mem.limits, n); return n }
	return g, mem, log
}
