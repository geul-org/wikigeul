//ff:func feature=dumpfile type=writer control=sequence
//ff:what TestWriteRun — the run appears whole under its name, no temporary left; an unwritable place is an error
package dumpfile

import (
	"math/rand"
	"os"
	"path/filepath"
	"slices"
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
	os.Symlink("/dev/full", filepath.Join(dir, "small.tmp"))
	if err := writeRun(filepath.Join(dir, "small"), []Record{{Kind: 'Q', Num: 1, Data: []byte("a")}}, 1); err == nil {
		t.Fatal("full disk at the close: want an error")
	}
	big := make([]byte, 1<<20)
	rand.New(rand.NewSource(1)).Read(big)
	recs := slices.Repeat([]Record{{Kind: 'Q', Num: 1, Data: big}}, 8)
	os.Symlink("/dev/full", filepath.Join(dir, "large.tmp"))
	if err := writeRun(filepath.Join(dir, "large"), recs, 1); err == nil {
		t.Fatal("full disk while writing: want an error")
	}
}
