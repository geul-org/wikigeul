//ff:func feature=cli type=command control=sequence
//ff:what dumpSampleJobs — the dump jobs a run trains its zstd dictionary on: the same Q entities the database path samples
package main

import (
	"github.com/geul-org/wikigeul/internal/dumpfile"
	"github.com/geul-org/wikigeul/pkg/pipeline"
)

// dumpSampleJobs gives dictSamples the entities it samples, with their dump
// lines: the first -samples Q entities of the ID list (one the dump lacks
// without a line, which buildAll skips as the database path does), else the
// first -samples Q entities the sort holds — the dump's items from Q<From> in
// number order, as wikipg export-geul lists them. A line that does not
// decode is no entity there (sampleJob). The slice is never nil,
// so dictSamples takes it as its list.
func dumpSampleJobs(c config, st dumpfile.State, ids []pipeline.PackJob) ([]pipeline.PackJob, error) {
	m, err := dumpfile.OpenMerger(dumpfile.Paths(c.SortDir, st.Runs))
	if err != nil {
		return nil, err
	}
	defer m.Close()
	out := []pipeline.PackJob{}
	send := func(j pipeline.PackJob) bool {
		s, ok := sampleJob(j, ids != nil)
		if j.Kind == 'Q' && ok && len(out) < c.Samples {
			out = append(out, s)
		}
		return j.Kind != 'Q' || len(out) < c.Samples
	}
	err = dumpJobs(newDumpCursor(m), pipeline.FirstQ(ids, c.Samples), send) // fills out: read it after
	return out, err
}
