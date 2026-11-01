//ff:func feature=dumpfile type=loader control=sequence
//ff:what TestOpenMerger — empty runs are dropped, the smallest front record is on top; a missing run is an error
package dumpfile

import (
	"path/filepath"
	"testing"
)

func TestOpenMerger(t *testing.T) {
	dir := t.TempDir()
	a, b, e := filepath.Join(dir, "a"), filepath.Join(dir, "b"), filepath.Join(dir, "e")
	writeRun(a, []Record{{Kind: 'Q', Num: 5}}, 1)
	writeRun(b, []Record{{Kind: 'P', Num: 7}}, 1)
	writeRun(e, nil, 1)
	m, err := OpenMerger([]string{a, e, b})
	if err != nil || len(m.h) != 2 || m.h[0].cur.Kind != 'P' {
		t.Fatalf("%v", err)
	}
	m.Close()
	if _, err := OpenMerger([]string{a, filepath.Join(dir, "none")}); err == nil {
		t.Fatal("want an error")
	}
}
