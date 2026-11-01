//ff:func feature=dumpfile type=loader control=iteration dimension=1
//ff:what TestMergerNext — records of all runs in order, then false; a run corrupt after its first record is an error
package dumpfile

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/klauspost/compress/zstd"
)

func TestMergerNext(t *testing.T) {
	dir := t.TempDir()
	a, b := filepath.Join(dir, "a"), filepath.Join(dir, "b")
	writeRun(a, []Record{{Kind: 'Q', Num: 1}, {Kind: 'Q', Num: 3}}, 1)
	writeRun(b, []Record{{Kind: 'Q', Num: 2}}, 1)
	m, _ := OpenMerger([]string{a, b})
	var got []uint64
	for r, ok, err := m.Next(); ok && err == nil; r, ok, err = m.Next() {
		got = append(got, r.Num)
	}
	if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Fatalf("%v", got)
	}
	z, _ := zstd.NewWriter(nil)
	bad := filepath.Join(dir, "bad")
	os.WriteFile(bad, z.EncodeAll(append(appendRecord(nil, Record{Kind: 'Q', Num: 1}), 'Q'), nil), 0o644)
	m, err := OpenMerger([]string{bad})
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err := m.Next(); err == nil {
		t.Fatal("want an error")
	}
	m.Close()
}
