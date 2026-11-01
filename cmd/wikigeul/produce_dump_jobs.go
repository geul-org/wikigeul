//ff:func feature=cli type=command control=sequence
//ff:what produceDumpJobs — send the run's entities of a sorted dump after the checkpoint, in key order, with their dump lines
package main

import (
	"github.com/geul-org/wikigeul/internal/dumpfile"
	"github.com/geul-org/wikigeul/pkg/pipeline"
)

// produceDumpJobs lists the jobs of a dump read directly: it merges the
// sort's runs (dumpfile.OpenMerger), numbers and sends the jobs on jobs
// (dumpJobs) and closes it. Entities up to the checkpoint's (p.Kind, p.Num)
// are passed over. An error is left in *errp.
func produceDumpJobs(c config, st dumpfile.State, p pipeline.PackProgress, ids []pipeline.PackJob, jobs chan<- pipeline.PackJob, errp *error) {
	defer close(jobs)
	m, err := dumpfile.OpenMerger(dumpfile.Paths(c.SortDir, st.Runs))
	if err != nil {
		*errp = err
		return
	}
	defer m.Close()
	var seq int64
	send := func(j pipeline.PackJob) bool {
		if pipeline.JobAfter(j.Kind, j.Num, p.Kind, p.Num) {
			j.Seq, seq = seq, seq+1
			jobs <- j
		}
		return true
	}
	*errp = dumpJobs(newDumpCursor(m), ids, send)
}
