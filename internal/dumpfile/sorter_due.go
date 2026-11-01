//ff:func feature=dumpfile type=helper control=sequence
//ff:what sorter.due — whether to write a run now: enough held for the memory there is, enough dump read, or the guard is stopping the sort
package dumpfile

// due asks for a run when the JSON held reaches c.Mem or a quarter of the
// process's memory budget, whichever is less (the records also take their
// slice, and the collector leaves room around them), when checkpointBytes of
// text were read, or when the guard stops the sort: the run empties memory
// and checkpoints. Run boundaries do not change the merged order.
func (s *sorter) due() bool {
	return s.held >= s.c.Guard.Share(s.c.Mem, sortBudgetPart) || s.since >= checkpointBytes || s.c.Guard.Err() != nil
}
