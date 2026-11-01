//ff:func feature=memguard type=loader control=sequence
//ff:what TestPeakRSS — VmHWM in bytes; 0 when the file or the line is missing
package memguard

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPeakRSS(t *testing.T) {
	p := filepath.Join(t.TempDir(), "status")
	os.WriteFile(p, []byte("Name: x\nVmHWM:   12 kB\n"), 0o644)
	if n := peakRSS(p); n != 12<<10 {
		t.Fatal(n)
	}
	if peakRSS(p+"x") != 0 {
		t.Fatal("no file")
	}
	os.WriteFile(p, []byte("Name: x\n"), 0o644)
	if peakRSS(p) != 0 {
		t.Fatal("no line")
	}
}
