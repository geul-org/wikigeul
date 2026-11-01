//ff:func feature=cli type=command control=sequence
//ff:what runDumpPack — pack mode over a sorted dump: the same stream, writer, workers and checkpoints, the jobs carrying their dump lines
package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/geul-org/wikigeul/internal/dumpfile"
	"github.com/geul-org/wikigeul/internal/port"
)

// runDumpPack is runPack for a dump sorted into st: the Issuer's pack set is
// opened (openPackStream: resumed, or new with the dictionary trained on
// dumpSampleJobs), the jobs come from the merged runs (produceDumpJobs) and
// c.Workers goroutines convert them with no database (loadJob decodes each
// line) into the same writer. -resume when the earlier run stopped while
// sorting, before the stream began, starts the stream.
func runDumpPack(c config, prov port.Provenance, st dumpfile.State, ids []packJob, idsHash string) error {
	set := filepath.Join(c.Root, ".well-known", "geul", "0x00010000")
	samples, err := dumpSampleJobs(c, st, ids)
	if err != nil {
		return fmt.Errorf("dictionary samples: %w", err)
	}
	if c.RecordedAt == 0 && !c.Resume {
		c.RecordedAt = time.Now().Unix()
	}
	s, p, err := openPackStream(set, c, prov, nil, samples, idsHash)
	if c.Resume && errors.Is(err, fs.ErrNotExist) {
		fmt.Printf("pack: no stream to resume in %s; starting it\n", set)
		c.Resume = false
		if c.RecordedAt == 0 {
			c.RecordedAt = time.Now().Unix()
		}
		s, p, err = openPackStream(set, c, prov, nil, samples, idsHash)
	}
	if err != nil {
		return err
	}
	c.RecordedAt = p.RecordedAt
	w, err := newPackWriter(s, c, p)
	if err != nil {
		return err
	}
	defer w.rep.Close()
	jobs := make(chan packJob, 4*c.Workers)
	var perr error
	go produceDumpJobs(c, st, p, ids, jobs, &perr)
	return w.drain(convertJobs(make([]*sql.DB, c.Workers), c, prov, s, jobs), &perr)
}
