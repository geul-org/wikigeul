//ff:func feature=pipeline type=command control=sequence
//ff:what WritePack — convert the run's jobs with parallel workers and stream them in key order into an opened pack stream, checkpointing, until the set closes
package pipeline

import (
	"github.com/geul-org/gogeul/pkg/pack"

	"github.com/geul-org/wikigeul/pkg/port"
)

// WritePack writes run c into stream s from where progress p stands: produce
// numbers the jobs still to convert, sends them and closes the channel,
// leaving an error in its second argument; one goroutine per loader converts
// them (convertJobs) and the writer puts the results in document key order
// (packWriter.take), each document as version 1 with its version list, served
// under h/ by content hash without a copy. The report is opened fresh or cut
// back to the checkpoint (newPackWriter).
func WritePack(s *pack.Stream, c Config, p PackProgress, prov port.Provenance, loaders []Loader, produce func(chan<- PackJob, *error)) error {
	c.RecordedAt = p.RecordedAt
	w, err := newPackWriter(s, c, p)
	if err != nil {
		return err
	}
	defer w.rep.Close()
	jobs := make(chan PackJob, 4*c.Workers)
	var perr error
	go produce(jobs, &perr)
	return w.drain(convertJobs(loaders, c, prov, s, jobs), &perr)
}
