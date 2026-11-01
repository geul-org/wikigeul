//ff:func feature=dumpfile type=writer control=sequence
//ff:what TestWriteRun — the run appears whole under its name, no temporary left; an unwritable place is an error
package dumpfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteRun(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "run")
	if err := writeRun(p, []Record{{Kind: 'Q', Num: 1, Data: []byte("a")}, {Kind: 'Q', Num: 2, Data: []byte("b")}}, 1); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(p + ".tmp"); err == nil {
		t.Fatal("temporary left")
	}
	m, _ := OpenMerger([]string{p})
	r1, _, _ := m.Next()
	r2, _, _ := m.Next()
	_, ok, _ := m.Next()
	if string(r1.Data) != "a" || string(r2.Data) != "b" || ok {
		t.Fatal("content")
	}
	if err := writeRun(filepath.Join(dir, "no", "run"), nil, 1); err == nil {
		t.Fatal("want an error")
	}
	os.Mkdir(filepath.Join(dir, "d"), 0o755)
	os.Mkdir(filepath.Join(dir, "d", "run"), 0o755)
	os.WriteFile(filepath.Join(dir, "d", "run", "x"), nil, 0o644)
	if err := writeRun(filepath.Join(dir, "d", "run"), nil, 1); err == nil {
		t.Fatal("rename onto a full directory: want an error")
	}
}
