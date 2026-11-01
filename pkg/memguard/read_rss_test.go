//ff:func feature=memguard type=loader control=sequence
//ff:what TestReadRSS — statm's second field in pages becomes bytes; no file, a short file and a non-number are errors; this process's is above 0
package memguard

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadRSS(t *testing.T) {
	p := filepath.Join(t.TempDir(), "statm")
	os.WriteFile(p, []byte("10 3 1\n"), 0o644)
	if n, err := readRSS(p); err != nil || n != 3*int64(os.Getpagesize()) {
		t.Fatalf("%d %v", n, err)
	}
	if _, err := readRSS(p + "x"); err == nil {
		t.Fatal("no file")
	}
	os.WriteFile(p, []byte("10\n"), 0o644)
	if _, err := readRSS(p); err == nil {
		t.Fatal("short")
	}
	os.WriteFile(p, []byte("10 x\n"), 0o644)
	if _, err := readRSS(p); err == nil {
		t.Fatal("not a number")
	}
	if n, err := readRSS(procStatm); err != nil || n <= 0 {
		t.Fatalf("own: %d %v", n, err)
	}
}
