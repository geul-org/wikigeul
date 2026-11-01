//ff:func feature=memguard type=loader control=sequence
//ff:what TestFindCgroup — the lowest limit from the process's cgroup up to the root, in v2 and v1; no self file and no limit give none
package memguard

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFindCgroup(t *testing.T) {
	root := t.TempDir()
	self := filepath.Join(root, "self")
	put := func(path, text string) {
		os.MkdirAll(filepath.Dir(path), 0o755)
		os.WriteFile(path, []byte(text), 0o644)
	}
	if c := findCgroup(self, root); c.Limit != 0 {
		t.Fatal("no self file")
	}
	put(self, "0::/a/b\n")
	put(filepath.Join(root, "a/b/memory.max"), "max\n")
	if c := findCgroup(self, root); c.Limit != 0 {
		t.Fatal("no limit")
	}
	put(filepath.Join(root, "a/memory.max"), "1000\n")
	put(filepath.Join(root, "memory.max"), "5000\n")
	if c := findCgroup(self, root); c.Limit != 1000 || !strings.HasSuffix(c.Current, "a/memory.current") {
		t.Fatalf("%+v", c)
	}
	put(filepath.Join(root, "a/b/memory.max"), "2000\n")
	if c := findCgroup(self, root); c.Limit != 1000 {
		t.Fatalf("%+v", c)
	}
	put(self, "4:memory:/x\n")
	put(filepath.Join(root, "memory/x/memory.limit_in_bytes"), "700\n")
	if c := findCgroup(self, root); c.Limit != 700 || c.Key != "total_inactive_file" {
		t.Fatalf("%+v", c)
	}
}
