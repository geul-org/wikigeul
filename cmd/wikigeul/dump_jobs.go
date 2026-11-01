//ff:func feature=cli type=helper control=selection
//ff:what dumpJobs — the jobs of a sorted dump in key order: the listed entities (missing ones without a line), or every entity
package main

// dumpJobs passes the run's jobs from the sorted dump d to send in document
// key order until send returns false: with an ID list, one per listed entity
// (listedJobs); else every entity of the sort (mergedJobs).
func dumpJobs(d *dumpCursor, ids []packJob, send func(packJob) bool) error {
	switch {
	case ids != nil:
		return listedJobs(d, ids, send)
	default:
		return mergedJobs(d, send)
	}
}
