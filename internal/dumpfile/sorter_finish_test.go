//ff:func feature=dumpfile type=writer control=sequence
//ff:what TestSorterFinish — the last run, the SHA-256 checked and the sort marked scanned; a wrong SHA-256, a failed spill or decompressor is an error
package dumpfile

import (
	"io"
	"os"
	"testing"
)

func TestSorterFinish(t *testing.T) {
	c := fixtureSortConfig(t, "d.json.gz", "[\n]\n")
	os.MkdirAll(c.Dir, 0o755)
	finish := func(c SortConfig, st *State) error {
		in, err := openInput(c, *st)
		if err != nil {
			return err
		}
		defer in.close()
		io.ReadAll(in.r)
		return (&sorter{c: c, st: st, in: in}).finish()
	}
	st := &State{SHA256: c.Info.SHA256}
	if err := finish(c, st); err != nil || !st.Scanned {
		t.Fatalf("%+v %v", st, err)
	}
	st = &State{SHA256: "00"}
	if err := finish(c, st); err == nil || st.Scanned {
		t.Fatal("wrong SHA-256: want an error")
	}
	bad := c
	bad.Dir = c.Dir + "/none"
	if err := finish(bad, &State{SHA256: c.Info.SHA256}); err == nil {
		t.Fatal("spill: want an error")
	}
	b := fixtureSortConfig(t, "d.json.bz2", "[\n]\n")
	os.MkdirAll(b.Dir, 0o755)
	os.WriteFile(b.Info.Path, []byte("not bzip2"), 0o644)
	if err := finish(b, &State{}); err == nil {
		t.Fatal("bzip2 failed: want an error")
	}
}
