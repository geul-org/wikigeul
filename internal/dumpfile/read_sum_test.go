//ff:func feature=dumpfile type=loader control=sequence
//ff:what TestReadSum — the hash of the line naming the file (binary marker and path allowed); no file or line is empty
package dumpfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadSum(t *testing.T) {
	p := filepath.Join(t.TempDir(), "sums")
	os.WriteFile(p, []byte("11  other.json\nbad line here x\n22 *dir/dump.json\n"), 0o644)
	if got := readSum(p, "dump.json"); got != "22" {
		t.Fatalf("got %q", got)
	}
	if got := readSum(p, "none.json"); got != "" {
		t.Fatalf("no line: %q", got)
	}
	if got := readSum(p+".missing", "dump.json"); got != "" {
		t.Fatalf("no file: %q", got)
	}
}
