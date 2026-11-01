//ff:func feature=cli type=command control=sequence
//ff:what runDumpPack — pack mode over a sorted dump: the shared stream, writer, workers and checkpoints, the jobs carrying their dump lines
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"
	"time"

	"github.com/geul-org/wikigeul/internal/dumpfile"
	"github.com/geul-org/wikigeul/pkg/pipeline"
	"github.com/geul-org/wikigeul/pkg/port"
)

// runDumpPack writes the dump sorted into st as the pack set of
// wikidata.geul.org under <root>/.well-known/geul/0x00010000: the Issuer's
// pack set is opened (pipeline.OpenPackStream: resumed, or new with the
// dictionary trained on dumpSampleJobs), the jobs come from the merged runs
// (produceDumpJobs) and c.Workers goroutines convert them with no database
// (loadJob decodes each line) into the same writer as the loaded database's
// (pipeline.WritePack). No documents.geul is written: at Wikidata's size the
// name index would be one 4 GB document, and a document's name follows from
// its ID (port.DocNumber). -resume when the earlier run stopped while
// sorting, before the stream began, starts the stream.
func runDumpPack(c config, prov port.Provenance, st dumpfile.State, ids []pipeline.PackJob, idsHash string) error {
	set := filepath.Join(c.Root, ".well-known", "geul", "0x00010000")
	jobs, err := dumpSampleJobs(c, st, ids)
	if err != nil {
		return fmt.Errorf("dictionary samples: %w", err)
	}
	samples := func() ([]pipeline.PackJob, error) { return jobs, nil }
	if c.RecordedAt == 0 && !c.Resume {
		c.RecordedAt = time.Now().Unix()
	}
	s, p, err := pipeline.OpenPackStream(set, c.Config, prov, loadJob, samples, idsHash)
	if c.Resume && errors.Is(err, fs.ErrNotExist) {
		fmt.Printf("pack: no stream to resume in %s; starting it\n", set)
		c.Resume = false
		if c.RecordedAt == 0 {
			c.RecordedAt = time.Now().Unix()
		}
		s, p, err = pipeline.OpenPackStream(set, c.Config, prov, loadJob, samples, idsHash)
	}
	if err != nil {
		return err
	}
	loaders := slices.Repeat([]pipeline.Loader{loadJob}, c.Workers)
	return pipeline.WritePack(s, c.Config, p, prov, loaders, func(out chan<- pipeline.PackJob, errp *error) {
		produceDumpJobs(c, st, p, ids, out, errp)
	})
}
