//ff:func feature=pipeline type=helper control=sequence
//ff:what fixtureGuard — a started guard seeing pct percent of memory free, stopped when the test ends; Confirm 1, so 10 percent is a stop at once
package pipeline

import (
	"io"
	"testing"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

// fixtureGuard starts a guard over a fixed reading of a 1 TiB machine.
func fixtureGuard(t *testing.T, pct int64) *memguard.Guard {
	t.Helper()
	read := func(string) (memguard.Sample, error) {
		return memguard.Sample{RSS: 1 << 30, Total: 1 << 40, Avail: (1 << 40) / 100 * pct}, nil
	}
	g, err := memguard.Start(memguard.Limits{Read: read, Confirm: 1}, 8, io.Discard)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(g.Close)
	return g
}
