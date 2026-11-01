//ff:func feature=memguard type=loader control=sequence
//ff:what TestCgroupAt — v2 takes the lower of max and high, v1 its limit, each with the files that measure it
package memguard

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCgroupAt(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "memory.max"), []byte("max\n"), 0o644)
	os.WriteFile(filepath.Join(dir, "memory.high"), []byte("900\n"), 0o644)
	os.WriteFile(filepath.Join(dir, "memory.limit_in_bytes"), []byte("800\n"), 0o644)
	if c := cgroupAt(dir, true); c.Limit != 900 || c.Key != "inactive_file" || c.Current != filepath.Join(dir, "memory.current") {
		t.Fatalf("%+v", c)
	}
	if c := cgroupAt(dir, false); c.Limit != 800 || c.Key != "total_inactive_file" || c.Current != filepath.Join(dir, "memory.usage_in_bytes") {
		t.Fatalf("%+v", c)
	}
}
