//ff:func feature=cli type=helper control=sequence
//ff:what TestMergedJobs — every sorted entity until send stops; a cursor ended by an error returns it
package main

import (
	"errors"
	"testing"

	"github.com/geul-org/wikigeul/internal/dumpfile"
	"github.com/geul-org/wikigeul/pkg/pipeline"
)

func TestMergedJobs(t *testing.T) {
	c, st := fixtureSortedDump(t, nil)
	m, _ := dumpfile.OpenMerger(dumpfile.Paths(c.SortDir, st.Runs))
	defer m.Close()
	var got []pipeline.PackJob
	err := mergedJobs(newDumpCursor(m), func(j pipeline.PackJob) bool { got = append(got, j); return len(got) < 2 })
	if err != nil || len(got) != 2 || got[0].Kind != 'P' || got[1].Num != 1 {
		t.Fatalf("%+v %v", got, err)
	}
	boom := errors.New("boom")
	if err := mergedJobs(&dumpCursor{err: boom}, func(pipeline.PackJob) bool { return true }); err != boom {
		t.Fatalf("%v", err)
	}
}
