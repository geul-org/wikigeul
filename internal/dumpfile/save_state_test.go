//ff:func feature=dumpfile type=writer control=sequence
//ff:what TestSaveState — the state is written whole and the temporary file is gone; a missing directory or a directory in the way is an error
package dumpfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveState(t *testing.T) {
	dir := t.TempDir()
	if err := saveState(dir, State{Next: 3}); err != nil {
		t.Fatal(err)
	}
	if st, ok, _ := loadState(dir); !ok || st.Next != 3 {
		t.Fatal("not saved")
	}
	if _, err := os.Stat(filepath.Join(dir, "state.json.tmp")); err == nil {
		t.Fatal("temporary left")
	}
	if err := saveState(filepath.Join(dir, "none"), State{}); err == nil {
		t.Fatal("no directory: want an error")
	}
	odd := t.TempDir()
	os.MkdirAll(filepath.Join(odd, "state.json", "x"), 0o755)
	if err := saveState(odd, State{}); err == nil {
		t.Fatal("directory in the way: want an error")
	}
}
