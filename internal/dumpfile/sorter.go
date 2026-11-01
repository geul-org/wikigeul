//ff:type feature=dumpfile type=model
//ff:what sorter — a sort in progress: config, state, the open dump and the records held for the next run
package dumpfile

// sorter reads one dump (scanDump): the records kept since the last run
// (recs, held bytes of JSON) and the dump bytes read since the last
// checkpoint (since).
type sorter struct {
	c     SortConfig
	st    *State
	in    *input
	recs  []Record
	held  int64
	since int64
}
