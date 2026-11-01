//ff:type feature=dumpfile type=loader
//ff:what Merger — the records of several sorted runs as one sorted stream
package dumpfile

// Merger merges sorted runs: h is a binary min-heap of the open runs by
// their front records, so Next yields every record of every run in run order
// (compareRecords) while holding one record per run.
type Merger struct {
	h []*runReader
}
