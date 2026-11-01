//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what listedJobs — a job per listed entity with its dump line, without one when the dump does not hold it
package main

// listedJobs sends a dump job for each entity of ids (in key order, readIDs)
// until send returns false: with its line when the sorted dump d holds it,
// without (loadJob: missing, as the database reports it) when not.
func listedJobs(d *dumpCursor, ids []packJob, send func(packJob) bool) error {
	for _, id := range ids {
		j := packJob{Kind: id.Kind, Num: id.Num, Dump: true}
		if d.seek(id.Kind, id.Num) {
			j = d.cur
		}
		if d.err != nil {
			return d.err
		}
		if !send(j) {
			return nil
		}
	}
	return nil
}
