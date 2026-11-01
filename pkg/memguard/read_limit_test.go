//ff:func feature=memguard type=loader control=sequence
//ff:what TestReadLimit — a number; 0 for no file, "max" and v1's no-limit value
package memguard

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadLimit(t *testing.T) {
	p := filepath.Join(t.TempDir(), "l")
	if readLimit(p) != 0 {
		t.Fatal("no file")
	}
	os.WriteFile(p, []byte("1234\n"), 0o644)
	if readLimit(p) != 1234 {
		t.Fatal("number")
	}
	os.WriteFile(p, []byte("max\n"), 0o644)
	if readLimit(p) != 0 {
		t.Fatal("max")
	}
	os.WriteFile(p, []byte("9223372036854771712\n"), 0o644)
	if readLimit(p) != 0 {
		t.Fatal("v1 no limit")
	}
}
