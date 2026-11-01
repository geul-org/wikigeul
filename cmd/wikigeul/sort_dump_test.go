//ff:func feature=cli type=command control=sequence
//ff:what TestSortDump — the range filter keeps every property and the Q range; a list keeps the listed; the filter names the sort
package main

import (
	"path/filepath"
	"testing"

	"github.com/geul-org/wikigeul/internal/dumpfile"
	"github.com/geul-org/wikigeul/pkg/pipeline"
)

func TestSortDump(t *testing.T) {
	c, st := fixtureSortedDump(t, nil)
	if st.Kept != 4 || st.Skipped != 1 || st.Filter != "P all, Q1-Q100" || !st.Done {
		t.Fatalf("%+v", st)
	}
	info, _ := dumpfile.InfoOf(c.DumpFile, "", "", 0)
	c.SortDir = filepath.Join(t.TempDir(), "l")
	st, err := sortDump(c, info, []pipeline.PackJob{{Kind: 'Q', Num: 42}}, "abc")
	if err != nil || st.Kept != 1 || st.Filter != "ids sha256:abc" {
		t.Fatalf("list %+v %v", st, err)
	}
}
