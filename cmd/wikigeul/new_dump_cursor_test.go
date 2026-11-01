//ff:func feature=cli type=helper control=sequence
//ff:what TestNewDumpCursor — a cursor on the first sorted entity (the property), as a dump job with its line
package main

import (
	"testing"

	"github.com/geul-org/wikigeul/internal/dumpfile"
)

func TestNewDumpCursor(t *testing.T) {
	c, st := fixtureSortedDump(t, nil)
	m, _ := dumpfile.OpenMerger(dumpfile.Paths(c.SortDir, st.Runs))
	defer m.Close()
	d := newDumpCursor(m)
	if !d.ok || d.err != nil || d.cur.Kind != 'P' || d.cur.Num != 31 || d.cur.Line == "" {
		t.Fatalf("%+v", d.cur)
	}
}
