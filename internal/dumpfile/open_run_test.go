//ff:func feature=dumpfile type=loader control=sequence
//ff:what TestOpenRun — a run opens at its first record; an empty run is false; a missing or corrupt file is an error
package dumpfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenRun(t *testing.T) {
	dir := t.TempDir()
	full, empty, bad := filepath.Join(dir, "full"), filepath.Join(dir, "empty"), filepath.Join(dir, "bad")
	writeRun(full, []Record{{Kind: 'Q', Num: 5, Data: []byte("x")}}, 1)
	writeRun(empty, nil, 1)
	os.WriteFile(bad, []byte("not zstd at all"), 0o644)
	rr, ok, err := openRun(full)
	if err != nil || !ok || rr.cur.Num != 5 {
		t.Fatalf("%v %v", ok, err)
	}
	rr.close()
	if _, ok, err := openRun(empty); ok || err != nil {
		t.Fatalf("empty: %v %v", ok, err)
	}
	if _, _, err := openRun(bad); err == nil {
		t.Fatal("corrupt: want an error")
	}
	if _, _, err := openRun(filepath.Join(dir, "none")); err == nil {
		t.Fatal("missing: want an error")
	}
}
