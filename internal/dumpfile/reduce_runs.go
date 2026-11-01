//ff:func feature=dumpfile type=writer control=iteration dimension=1
//ff:what reduceRuns — merge runs mergeFanIn at a time until at most mergeFanIn are left, then mark the sort done
package dumpfile

// reduceRuns merges the first mergeFanIn runs into one (mergeFirst) while
// there are more than mergeFanIn, so the final merge holds a bounded number
// of decoders, then marks the sort done.
func reduceRuns(c SortConfig, st *State) error {
	for len(st.Runs) > mergeFanIn {
		if err := mergeFirst(c, st); err != nil {
			return err
		}
	}
	st.Done = true
	return saveState(c.Dir, *st)
}
