//ff:func feature=cli type=helper control=sequence
//ff:what fixtureSortedDump — a pack-mode config over testDumpBody sorted into its sort directory, with the sort state
package main

import (
	"path/filepath"
	"testing"

	"github.com/geul-org/wikigeul/internal/dumpfile"
)

// fixtureSortedDump sorts testDumpBody for c (every property and Q1–Q100, or
// ids) and returns c with the dump fields set and the sort state.
func fixtureSortedDump(t *testing.T, ids []packJob) (config, dumpfile.State) {
	t.Helper()
	c := packTestConfig(t, "wikidata")
	c.DumpFile, c.SortDir, c.SortMem, c.RefOrder = fixtureDumpFile(t, testDumpBody), filepath.Join(t.TempDir(), "sort"), 1, true
	info, err := dumpfile.InfoOf(c.DumpFile, "", "", 0)
	if err != nil {
		t.Fatal(err)
	}
	c.Source = info.Name + " sha256:" + info.SHA256
	st, err := sortDump(c, info, ids, "h")
	if err != nil {
		t.Fatal(err)
	}
	return c, st
}
