//ff:func feature=memguard type=loader control=sequence
//ff:what TestProcReader — this system's numbers read; a meminfo path of its own is used; a missing one fails
package memguard

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProcReader(t *testing.T) {
	s, err := procReader("")("x")
	if err != nil || s.Total <= 0 || s.Avail < 0 || s.RSS <= 0 {
		t.Fatalf("%+v %v", s, err)
	}
	p := filepath.Join(t.TempDir(), "meminfo")
	os.WriteFile(p, []byte("MemTotal: 8 kB\nMemAvailable: 7 kB\n"), 0o644)
	if s, err = procReader(p)("x"); err != nil || s.Avail != 7<<10 {
		t.Fatalf("%+v %v", s, err)
	}
	if _, err = procReader(p + "x")("x"); err == nil {
		t.Fatal("want an error")
	}
}
