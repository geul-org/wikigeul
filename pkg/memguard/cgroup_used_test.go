//ff:func feature=memguard type=loader control=sequence
//ff:what TestCgroupUsed — current less the inactive file line, never below 0; no such line keeps current; unreadable or bad files are errors
package memguard

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCgroupUsed(t *testing.T) {
	dir := t.TempDir()
	c := cgroup{Current: filepath.Join(dir, "cur"), Stat: filepath.Join(dir, "stat"), Key: "inactive_file"}
	if _, err := c.used(); err == nil {
		t.Fatal("no current")
	}
	os.WriteFile(c.Current, []byte("x"), 0o644)
	if _, err := c.used(); err == nil {
		t.Fatal("bad current")
	}
	os.WriteFile(c.Current, []byte("100\n"), 0o644)
	if _, err := c.used(); err == nil {
		t.Fatal("no stat")
	}
	os.WriteFile(c.Stat, []byte("anon 5\ninactive_file 30\n"), 0o644)
	if n, err := c.used(); err != nil || n != 70 {
		t.Fatalf("%d %v", n, err)
	}
	os.WriteFile(c.Stat, []byte("inactive_file 300\n"), 0o644)
	if n, _ := c.used(); n != 0 {
		t.Fatal(n)
	}
	os.WriteFile(c.Stat, []byte("anon 5\n"), 0o644)
	if n, _ := c.used(); n != 100 {
		t.Fatal(n)
	}
}
