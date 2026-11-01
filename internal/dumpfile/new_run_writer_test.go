//ff:func feature=dumpfile type=writer control=sequence
//ff:what TestNewRunWriter — a writer over a new file; a path in no directory is an error
package dumpfile

import (
	"path/filepath"
	"testing"
)

func TestNewRunWriter(t *testing.T) {
	w, err := newRunWriter(filepath.Join(t.TempDir(), "r"), 0)
	if err != nil || w.f == nil || w.z == nil {
		t.Fatalf("%v", err)
	}
	w.close()
	if _, err := newRunWriter(filepath.Join(t.TempDir(), "no", "r"), 1); err == nil {
		t.Fatal("want an error")
	}
}
