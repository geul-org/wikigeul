//ff:func feature=dumpfile type=writer control=sequence
//ff:what mergeFirst — merge the first n runs into a new run at the end, save the state, remove the merged runs
package dumpfile

import (
	"fmt"
	"path/filepath"
	"slices"
)

// mergeFirst merges the first n runs of st into the run numbered
// st.Next (mergeRuns), which goes to the end of the list, saves the state
// and only then removes the merged files. Interrupted before the save, by a
// crash or the guard, the same merge is done again into the same name.
func mergeFirst(c SortConfig, st *State, n int) error {
	ins := slices.Clone(st.Runs[:n])
	name := fmt.Sprintf("run-%06d.zst", st.Next)
	if err := mergeRuns(filepath.Join(c.Dir, name), Paths(c.Dir, ins), c.Guard.Active(c.Threads), c.Guard); err != nil {
		return err
	}
	st.Runs = append(slices.Clone(st.Runs[n:]), name)
	st.Next++
	if err := saveState(c.Dir, *st); err != nil {
		return err
	}
	fmt.Fprintf(c.Log, "dump sort: merged %d runs into %s, %d runs left\n", len(ins), name, len(st.Runs))
	removeFiles(Paths(c.Dir, ins))
	return nil
}
