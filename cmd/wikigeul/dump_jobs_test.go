//ff:func feature=cli type=helper control=sequence
//ff:what TestDumpJobs — with a list only the listed entities (a missing one without a line), without one every entity
package main

import (
	"testing"

	"github.com/geul-org/wikigeul/internal/dumpfile"
	"github.com/geul-org/wikigeul/pkg/pipeline"
)

func TestDumpJobs(t *testing.T) {
	c, st := fixtureSortedDump(t, nil)
	collect := func(ids []pipeline.PackJob) []pipeline.PackJob {
		m, _ := dumpfile.OpenMerger(dumpfile.Paths(c.SortDir, st.Runs))
		defer m.Close()
		var got []pipeline.PackJob
		if err := dumpJobs(newDumpCursor(m), ids, func(j pipeline.PackJob) bool { got = append(got, j); return true }); err != nil {
			t.Fatal(err)
		}
		return got
	}
	if got := collect([]pipeline.PackJob{{Kind: 'Q', Num: 2}, {Kind: 'Q', Num: 42}}); len(got) != 2 || got[0].Line != "" || got[1].Line == "" {
		t.Fatalf("listed %+v", got)
	}
	if got := collect(nil); len(got) != 4 {
		t.Fatalf("all %+v", got)
	}
}
