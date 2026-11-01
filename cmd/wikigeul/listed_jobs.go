//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what listedJobs — a job per listed entity with its dump line, without one when the dump does not hold it
package main

import "github.com/geul-org/wikigeul/pkg/pipeline"

// listedJobs sends a dump job for each entity of ids (in key order, ReadIDs)
// until send returns false: with its line when the sorted dump d holds it,
// without (loadJob: missing, as the database reports it) when not.
func listedJobs(d *dumpCursor, ids []pipeline.PackJob, send func(pipeline.PackJob) bool) error {
	for _, id := range ids {
		j := pipeline.PackJob{Kind: id.Kind, Num: id.Num}
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
