//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what mergedJobs — every entity of a sorted dump as a job, in key order
package main

// mergedJobs sends every entity of d until send returns false or the dump
// ends, and returns the error that ended it early.
func mergedJobs(d *dumpCursor, send func(packJob) bool) error {
	for d.ok {
		if !send(d.cur) {
			return nil
		}
		d.advance()
	}
	return d.err
}
