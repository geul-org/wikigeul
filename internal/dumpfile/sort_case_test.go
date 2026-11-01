//ff:func feature=dumpfile type=command control=iteration dimension=1
//ff:what sortCase — sort one test dump file: counts, key order, a second Sort returning the finished state, another SHA-256 refused
package dumpfile

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"testing"
)

// testDump has, in no ID order, two items in range, a property, a Lexeme
// (skipped), an ID with a leading zero (not read) and an item out of range.
const testDump = "[\n" +
	`{"type":"item","id":"Q5","labels":{}},` + "\n" +
	`{"type":"property","id":"P31","datatype":"wikibase-item"},` + "\n" +
	`{"type":"item","id":"Q1"},` + "\n" +
	`{"type":"lexeme","id":"L7"},` + "\n" +
	`{"type":"item","id":"Q01"},` + "\n" +
	`{"type":"item","id":"Q9000"}` + "\n" +
	"]\n"

// sortCase writes body as the dump name in a temporary directory and sorts
// it with 1 byte per run (a run per kept line, merged in the end).
func sortCase(t *testing.T, name string, body []byte) {
	dir := t.TempDir()
	path := filepath.Join(dir, name)
	os.WriteFile(path, body, 0o644)
	sum := sha256.Sum256(body)
	info, err := InfoOf(path, "", hex.EncodeToString(sum[:]), 0)
	if err != nil || info.URL != "https://dumps.wikimedia.org/wikidatawiki/entities/20261019/"+name {
		t.Fatalf("%s: %+v %v", name, info, err)
	}
	c := SortConfig{Info: info, Dir: filepath.Join(dir, "sort"), Mem: 1, Threads: 1, Filter: "f", Log: io.Discard,
		Keep: func(k byte, n uint64) bool { return k == 'P' || k == 'Q' && n <= 1000 }}
	st, err := Sort(c)
	if err != nil || !st.Done || st.Kept != 3 || st.Skipped != 2 || st.Bad != 1 || st.Lines != 8 {
		t.Fatalf("%s: %+v %v", name, st, err)
	}
	m, _ := OpenMerger(Paths(c.Dir, st.Runs))
	var got []byte
	for r, ok, _ := m.Next(); ok; r, ok, _ = m.Next() {
		got = append(got, r.Kind)
	}
	m.Close()
	if string(got) != "PQQ" {
		t.Errorf("%s: order %s", name, got)
	}
	if again, err := Sort(c); err != nil || len(again.Runs) != len(st.Runs) {
		t.Errorf("%s: again %+v %v", name, again, err)
	}
	c.Info.SHA256 = "00"
	if _, err := Sort(c); err == nil {
		t.Errorf("%s: another SHA-256 in the same directory: want an error", name)
	}
}
