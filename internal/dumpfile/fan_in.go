//ff:func feature=dumpfile type=helper control=sequence
//ff:what fanIn — how many runs a merge reads at once: mergeFanIn, fewer when the guard has less memory to give
package dumpfile

// fanIn fits the open runs to the memory above the guard's caution floor.
func fanIn(c SortConfig) int {
	return c.Guard.Fit(openRunMemory, minFanIn, mergeFanIn)
}
