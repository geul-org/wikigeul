//ff:func feature=cli type=helper control=sequence
//ff:what TestDumpCursorSeek — on an entity the dump holds, false for one it lacks (staying on the next), false past the end
package main

import (
	"testing"

	"github.com/geul-org/wikigeul/internal/dumpfile"
)

func TestDumpCursorSeek(t *testing.T) {
	c, st := fixtureSortedDump(t, nil)
	m, _ := dumpfile.OpenMerger(dumpfile.Paths(c.SortDir, st.Runs))
	defer m.Close()
	d := newDumpCursor(m)
	if !d.seek('P', 31) || d.cur.Num != 31 {
		t.Fatal("P31")
	}
	if d.seek('Q', 5) || d.cur.Num != 7 {
		t.Fatalf("Q5: on %c%d", d.cur.Kind, d.cur.Num)
	}
	if !d.seek('Q', 42) {
		t.Fatal("Q42")
	}
	if d.seek('Q', 100) || d.ok {
		t.Fatal("past the end")
	}
}
