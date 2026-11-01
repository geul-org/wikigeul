//ff:func feature=dumpfile type=writer control=sequence
//ff:what TestRemoveFiles — files go, missing ones are no error
package dumpfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveFiles(t *testing.T) {
	p := filepath.Join(t.TempDir(), "a")
	os.WriteFile(p, nil, 0o644)
	removeFiles([]string{p, p + ".none"})
	if _, err := os.Stat(p); err == nil {
		t.Fatal("still there")
	}
}
