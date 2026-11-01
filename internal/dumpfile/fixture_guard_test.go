//ff:func feature=dumpfile type=helper control=sequence
//ff:what fixtureGuard — a started guard over a fixed sample, Confirm 1, closed when the test ends
package dumpfile

import (
	"io"
	"testing"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

// fixtureGuard starts a guard for workers workers seeing s.
func fixtureGuard(t *testing.T, s memguard.Sample, workers int) *memguard.Guard {
	t.Helper()
	read := func(string) (memguard.Sample, error) { return s, nil }
	g, err := memguard.Start(memguard.Limits{Read: read, Confirm: 1}, workers, io.Discard)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(g.Close)
	return g
}
