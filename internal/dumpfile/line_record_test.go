//ff:func feature=dumpfile type=loader control=sequence
//ff:what TestLineRecord — no entity, kept with its line, skipped by the filter, unreadable id or ID that does not split
package dumpfile

import "testing"

func TestLineRecord(t *testing.T) {
	keep := func(k byte, n uint64) bool { return k == 'Q' }
	if _, s, err := lineRecord([]byte("[\n"), 1, keep); s != lineNone || err != nil {
		t.Fatal("bracket")
	}
	if r, s, err := lineRecord([]byte("{\"id\":\"Q4\"},\n"), 2, keep); s != lineKept || err != nil || r.Kind != 'Q' || r.Num != 4 || r.Line != 2 || string(r.Data) != `{"id":"Q4"}` {
		t.Fatalf("kept %+v %v %v", r, s, err)
	}
	if _, s, err := lineRecord([]byte(`{"id":"P4"}`), 3, keep); s != lineSkipped || err != nil {
		t.Fatal("skipped")
	}
	if _, s, err := lineRecord([]byte(`{"type":"item"}`), 4, keep); s != lineBad || err == nil {
		t.Fatal("no id")
	}
	if _, s, err := lineRecord([]byte(`{"id":"Q04"}`), 5, keep); s != lineBad || err == nil {
		t.Fatal("bad ID")
	}
}
