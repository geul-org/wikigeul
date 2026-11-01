//ff:func feature=dumpfile type=writer control=iteration dimension=1
//ff:what TestReduceRuns — more than mergeFanIn runs merge down to at most mergeFanIn keeping every record; a missing run is an error
package dumpfile

import (
	"fmt"
	"io"
	"path/filepath"
	"testing"
)

func TestReduceRuns(t *testing.T) {
	dir := t.TempDir()
	st := &State{}
	for i := 0; i < mergeFanIn+2; i++ {
		name := fmt.Sprintf("run-%06d.zst", i)
		writeRun(filepath.Join(dir, name), []Record{{Kind: 'Q', Num: uint64(i)}}, 1)
		st.Runs, st.Next = append(st.Runs, name), i+1
	}
	c := SortConfig{Dir: dir, Threads: 1, Log: io.Discard}
	if err := reduceRuns(c, st); err != nil || !st.Done || len(st.Runs) != 3 {
		t.Fatalf("%d runs %v", len(st.Runs), err)
	}
	m, _ := OpenMerger(Paths(dir, st.Runs))
	n := 0
	for _, ok, _ := m.Next(); ok; _, ok, _ = m.Next() {
		n++
	}
	m.Close()
	if n != mergeFanIn+2 {
		t.Fatalf("%d records", n)
	}
	st.Done, st.Runs = false, append(st.Runs, make([]string, mergeFanIn)...)
	if err := reduceRuns(c, st); err == nil {
		t.Fatal("missing runs: want an error")
	}
}
