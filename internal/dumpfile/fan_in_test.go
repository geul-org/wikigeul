//ff:func feature=dumpfile type=helper control=sequence
//ff:what TestFanIn — mergeFanIn with no guard or plenty; narrow when memory stands at the caution floor
package dumpfile

import (
	"testing"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

func TestFanIn(t *testing.T) {
	if fanIn(SortConfig{}) != mergeFanIn {
		t.Fatal("no guard")
	}
	s := memguard.Sample{RSS: 1 << 30, Total: 1 << 40, Avail: 1 << 39}
	if got := fanIn(SortConfig{Guard: fixtureGuard(t, s, 2)}); got != mergeFanIn {
		t.Fatal(got)
	}
	s.Avail = 1 << 38
	if got := fanIn(SortConfig{Guard: fixtureGuard(t, s, 2)}); got != minFanIn {
		t.Fatal(got)
	}
}
