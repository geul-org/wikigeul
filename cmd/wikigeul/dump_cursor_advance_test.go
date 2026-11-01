//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestDumpCursorAdvance — entities in key order, a repeated ID once (its first line), then the end
package main

import (
	"path/filepath"
	"testing"

	"github.com/geul-org/wikigeul/internal/dumpfile"
)

func TestDumpCursorAdvance(t *testing.T) {
	c, st := fixtureSortedDump(t, nil)
	m, _ := dumpfile.OpenMerger(dumpfile.Paths(c.SortDir, st.Runs))
	d := newDumpCursor(m)
	var got []string
	for ; d.ok; d.advance() {
		got = append(got, string(d.cur.Kind)+string(rune('0'+d.cur.Num%10)))
	}
	m.Close()
	if len(got) != 4 || got[0] != "P1" || got[1] != "Q1" || got[2] != "Q7" || got[3] != "Q2" || d.err != nil {
		t.Fatalf("%v %v", got, d.err)
	}
	c2 := c
	c2.DumpFile = fixtureDumpFile(t, "[\n{\"id\":\"Q5\",\"x\":1},\n{\"id\":\"Q5\",\"x\":2}\n]\n")
	info, _ := dumpfile.InfoOf(c2.DumpFile, "", "", 0)
	c2.SortDir = filepath.Join(t.TempDir(), "s")
	st2, err := sortDump(c2, info, nil, "")
	if err != nil {
		t.Fatal(err)
	}
	m, _ = dumpfile.OpenMerger(dumpfile.Paths(c2.SortDir, st2.Runs))
	d = newDumpCursor(m)
	first := d.cur.Line
	d.advance()
	m.Close()
	if first != `{"id":"Q5","x":1}` || d.ok || d.err != nil {
		t.Fatalf("repeated: %q %v %v", first, d.ok, d.err)
	}
}
