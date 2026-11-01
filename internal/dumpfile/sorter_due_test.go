//ff:func feature=dumpfile type=helper control=sequence
//ff:what TestSorterDue — due at c.Mem held, or checkpointBytes read; with a guard, at its share of the budget and at once when it stops
package dumpfile

import (
	"testing"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

func TestSorterDue(t *testing.T) {
	s := &sorter{c: SortConfig{Mem: 1600}}
	if s.due() {
		t.Fatal("empty")
	}
	s.held = 1600
	if !s.due() {
		t.Fatal("full")
	}
	s.held, s.since = 0, checkpointBytes
	if !s.due() {
		t.Fatal("text read")
	}
	s.since, s.held = 0, 100
	s.c.Guard = fixtureGuard(t, memguard.Sample{RSS: 400 * sortBudgetPart, Total: 1 << 40, Avail: 3 << 38}, 8)
	if s.due() {
		t.Fatal("plenty")
	}
	s.c.Guard = fixtureGuard(t, memguard.Sample{RSS: 1 << 30, Total: 1 << 40, Avail: 1 << 36}, 1)
	if !s.due() {
		t.Fatal("stopping")
	}
}
