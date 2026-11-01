//ff:func feature=dumpfile type=helper control=sequence
//ff:what TestBzip2Program — lbzip2 when PATH has it, else bzip2
package dumpfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBzip2Program(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("PATH", dir)
	if got := bzip2Program(); got != "bzip2" {
		t.Fatalf("got %s", got)
	}
	os.WriteFile(filepath.Join(dir, "lbzip2"), []byte("#!/bin/sh\n"), 0o755)
	if got := bzip2Program(); got != "lbzip2" {
		t.Fatalf("got %s", got)
	}
}
