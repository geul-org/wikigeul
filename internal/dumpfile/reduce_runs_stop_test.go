//ff:func feature=dumpfile type=writer control=iteration dimension=1
//ff:what TestReduceRunsStop — a guard stop ends the merging with the sort not done, and a merge in progress leaves no run named; memory short narrows the merges; a sort done stays as it is
package dumpfile

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"testing"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

func TestReduceRunsStop(t *testing.T) {
	dir := t.TempDir()
	st := &State{}
	for i := 0; i < 20; i++ {
		name := fmt.Sprintf("run-%06d.zst", i)
		writeRun(filepath.Join(dir, name), []Record{{Kind: 'Q', Num: uint64(i)}}, 1)
		st.Runs, st.Next = append(st.Runs, name), i+1
	}
	short := fixtureGuard(t, memguard.Sample{RSS: 1, Total: 1 << 40, Avail: 1 << 38}, 2)
	if got := fanIn(SortConfig{Guard: short}); got != minFanIn {
		t.Fatal(got)
	}
	c := SortConfig{Dir: dir, Threads: 1, Log: io.Discard, Guard: fixtureGuard(t, memguard.Sample{RSS: 1, Total: 1 << 40, Avail: 1 << 36}, 1)}
	var stop memguard.StopError
	if err := reduceRuns(c, st); !errors.As(err, &stop) || st.Done || len(st.Runs) != 20 {
		t.Fatalf("%v %+v", err, st)
	}
	if err := mergeRuns(filepath.Join(dir, "m"), Paths(dir, st.Runs[:2]), 1, c.Guard); !errors.As(err, &stop) {
		t.Fatal(err)
	}
	c.Guard = short
	if err := reduceRuns(c, st); err != nil || !st.Done || len(st.Runs) > minFanIn {
		t.Fatalf("%v %+v", err, st)
	}
	st.Runs = append(st.Runs, make([]string, 2*minFanIn)...)
	if err := reduceRuns(c, st); err != nil {
		t.Fatalf("done sort: %v", err)
	}
}
