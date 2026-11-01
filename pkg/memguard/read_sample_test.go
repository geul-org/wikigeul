//ff:func feature=memguard type=loader control=sequence
//ff:what TestReadSample — the numbers from the files, with the cgroup's when it has a limit and can be read; meminfo or statm unreadable is an error
package memguard

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadSample(t *testing.T) {
	dir := t.TempDir()
	mi, st, cur, stat := filepath.Join(dir, "mi"), filepath.Join(dir, "st"), filepath.Join(dir, "cur"), filepath.Join(dir, "stat")
	os.WriteFile(mi, []byte("MemTotal: 4 kB\nMemAvailable: 2 kB\n"), 0o644)
	os.WriteFile(st, []byte("1 2\n"), 0o644)
	os.WriteFile(cur, []byte("500\n"), 0o644)
	os.WriteFile(stat, []byte("inactive_file 100\n"), 0o644)
	cg := cgroup{Limit: 1000, Current: cur, Stat: stat, Key: "inactive_file"}
	s, err := readSample(mi, st, cg)
	if err != nil || s.Total != 4<<10 || s.Avail != 2<<10 || s.RSS != 2*int64(os.Getpagesize()) || s.CgroupLimit != 1000 || s.CgroupUsed != 400 {
		t.Fatalf("%+v %v", s, err)
	}
	if s, err = readSample(mi, st, cgroup{}); err != nil || s.CgroupLimit != 0 {
		t.Fatalf("no cgroup: %+v %v", s, err)
	}
	if _, err = readSample(mi+"x", st, cg); err == nil {
		t.Fatal("no meminfo")
	}
	if _, err = readSample(st, st, cg); err == nil {
		t.Fatal("bad meminfo")
	}
	if _, err = readSample(mi, st+"x", cg); err == nil {
		t.Fatal("no statm")
	}
}
