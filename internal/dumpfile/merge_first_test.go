//ff:func feature=dumpfile type=writer control=iteration dimension=1
//ff:what TestMergeFirst — the first mergeFanIn runs become one at the end, saved, the inputs removed; a failed merge or save is an error
package dumpfile

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestMergeFirst(t *testing.T) {
	dir := t.TempDir()
	st := &State{}
	for i := 0; i < mergeFanIn+1; i++ {
		name := fmt.Sprintf("run-%06d.zst", i)
		writeRun(filepath.Join(dir, name), []Record{{Kind: 'Q', Num: uint64(i)}}, 1)
		st.Runs, st.Next = append(st.Runs, name), i+1
	}
	first := filepath.Join(dir, st.Runs[0])
	c := SortConfig{Dir: dir, Threads: 1, Log: io.Discard}
	if err := mergeFirst(c, st); err != nil || len(st.Runs) != 2 || st.Runs[1] != fmt.Sprintf("run-%06d.zst", mergeFanIn+1) {
		t.Fatalf("%v %v", st.Runs, err)
	}
	if _, err := os.Stat(first); err == nil {
		t.Fatal("input left")
	}
	if saved, _, _ := loadState(dir); len(saved.Runs) != 2 {
		t.Fatal("not saved")
	}
	st.Runs = make([]string, mergeFanIn)
	if err := mergeFirst(c, st); err == nil {
		t.Fatal("missing runs: want an error")
	}
	ro := t.TempDir()
	st = &State{}
	for i := 0; i < mergeFanIn; i++ {
		name := fmt.Sprintf("r%d", i)
		writeRun(filepath.Join(ro, name), nil, 1)
		st.Runs = append(st.Runs, name)
	}
	os.MkdirAll(filepath.Join(ro, "state.json", "x"), 0o755)
	if err := mergeFirst(SortConfig{Dir: ro, Threads: 1, Log: io.Discard}, st); err == nil {
		t.Fatal("state not saved: want an error")
	}
}
