//ff:func feature=dumpfile type=writer control=iteration dimension=1
//ff:what reduceRuns — merge runs fanIn at a time until at most fanIn are left, then mark the sort done
package dumpfile

// reduceRuns merges the first fanIn(c) runs into one (mergeFirst) while there
// are more than that many, so the final merge holds a bounded number of
// decoders, then marks the sort done. The width follows the memory there is
// and is asked again for each merge; a finished sort is left alone.
func reduceRuns(c SortConfig, st *State) error {
	c.Guard.SetPhase("sort:merge")
	for w := fanIn(c); !st.Done && len(st.Runs) > w; w = fanIn(c) {
		if err := c.Guard.Err(); err != nil {
			return err
		}
		if err := mergeFirst(c, st, w); err != nil {
			return err
		}
	}
	st.Done = true
	return saveState(c.Dir, *st)
}
