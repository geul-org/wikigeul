//ff:func feature=memguard type=command control=sequence
//ff:what TestStart — off is a nil guard; memory that cannot be read fails; a guard on this system has sampled before it returns and closes
package memguard

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStart(t *testing.T) {
	if g, err := Start(Limits{Off: true}, 2, io.Discard); g != nil || err != nil {
		t.Fatalf("%v %v", g, err)
	}
	bad := func(string) (Sample, error) { return Sample{}, errors.New("no") }
	if _, err := Start(Limits{Read: bad}, 2, io.Discard); err == nil || !strings.Contains(err.Error(), "-mem-off") {
		t.Fatalf("%v", err)
	}
	p := filepath.Join(t.TempDir(), "meminfo")
	os.WriteFile(p, []byte("MemTotal: 1000000 kB\nMemAvailable: 900000 kB\n"), 0o644)
	g, err := Start(Limits{Meminfo: p}, 2, io.Discard)
	if err != nil {
		t.Fatal(err)
	}
	defer g.Close()
	if g.Level() != Normal || g.budget.Load() <= 0 {
		t.Fatalf("level %v budget %d", g.Level(), g.budget.Load())
	}
}
