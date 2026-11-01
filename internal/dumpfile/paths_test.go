//ff:func feature=dumpfile type=helper control=sequence
//ff:what TestPaths — run names joined to the directory, in order
package dumpfile

import (
	"path/filepath"
	"testing"
)

func TestPaths(t *testing.T) {
	got := Paths("d", []string{"a", "b"})
	if len(got) != 2 || got[0] != filepath.Join("d", "a") || got[1] != filepath.Join("d", "b") || len(Paths("d", nil)) != 0 {
		t.Fatalf("%v", got)
	}
}
