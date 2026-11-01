//ff:func feature=dumpfile type=command control=iteration dimension=1
//ff:what TestScanDump — runs written as memory fills, the dump read to its end and checked; resumed from a saved offset it gives the same records; a read failure or missing file is an error
package dumpfile

import (
	"os"
	"testing"
)

func TestScanDump(t *testing.T) {
	const body = "[\n{\"id\":\"Q3\"},\n{\"id\":\"Q1\"},\n{\"id\":\"P2\"}\n]\n"
	for _, name := range []string{"d.json", "d.json.gz"} {
		c := fixtureSortConfig(t, name, body)
		c.Mem = 1
		os.MkdirAll(c.Dir, 0o755)
		st := &State{SHA256: c.Info.SHA256}
		if err := scanDump(c, st); err != nil || !st.Scanned || len(st.Runs) != 3 || st.Kept != 3 || st.Lines != 5 {
			t.Fatalf("%s: %+v %v", name, st, err)
		}
		half := &State{SHA256: c.Info.SHA256}
		saved, _, _ := loadState(c.Dir)
		*half = saved
		half.Scanned, half.Runs, half.Kept = false, half.Runs[:1], 1
		half.Offset, half.Lines = int64(len("[\n{\"id\":\"Q3\"},\n")), 2
		if name == "d.json" {
			half.Hash = nil
			half.Offset = 0
			half.Lines = 0
			half.Runs, half.Kept = nil, 0
		}
		if err := scanDump(c, half); err != nil || half.Kept != 3 {
			t.Fatalf("%s resumed: %+v %v", name, half, err)
		}
	}
	nodir := fixtureSortConfig(t, "d.json", body)
	nodir.Mem = 1
	if err := scanDump(nodir, &State{}); err == nil {
		t.Fatal("run not writable: want an error")
	}
	c := fixtureSortConfig(t, "d.json.bz2", body)
	os.MkdirAll(c.Dir, 0o755)
	os.WriteFile(c.Info.Path, []byte("BZh9 broken"), 0o644)
	if err := scanDump(c, &State{}); err == nil {
		t.Fatal("broken bzip2: want an error")
	}
	c.Info.Path += ".none"
	if err := scanDump(c, &State{}); err == nil {
		t.Fatal("missing: want an error")
	}
}
