//ff:func feature=dumpfile type=loader control=sequence
//ff:what TestMergerClose — every open run is closed and the merger is empty
package dumpfile

import (
	"path/filepath"
	"testing"
)

func TestMergerClose(t *testing.T) {
	p := filepath.Join(t.TempDir(), "a")
	writeRun(p, []Record{{Kind: 'Q', Num: 1}}, 1)
	m, _ := OpenMerger([]string{p})
	f := m.h[0].f
	m.Close()
	if _, err := f.Stat(); err == nil || m.h != nil {
		t.Fatal("not closed")
	}
	if _, ok, _ := m.Next(); ok {
		t.Fatal("records after close")
	}
}
