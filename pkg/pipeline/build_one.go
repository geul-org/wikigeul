//ff:func feature=pipeline type=command control=sequence
//ff:what buildOne — load one entity (database or dump line, by its Loader) and build its document with fresh LOCAL IDs (pack mode); false when the dump lacks it
package pipeline

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/delta"

	"github.com/geul-org/wikigeul/pkg/port"
)

// buildOne loads entity job (load: the database, or its dump line) and builds its document
// (port.BuildDoc) with prov and the run's registry. A pack-mode run writes a
// new set, so no earlier version lends its IDs: they start from 1, as for a
// document never published in directory mode.
func buildOne(load Loader, c Config, prov port.Provenance, job PackJob) ([]byte, port.Report, bool, error) {
	id := string(job.Kind) + strconv.FormatUint(job.Num, 10)
	e, ok, err := load(job)
	if err != nil || !ok {
		return nil, port.Report{ID: id, Missing: err == nil}, false, err
	}
	ids, _ := delta.NewLocalIDs(nil, 0) // no previous version: never fails
	b, rep, err := port.BuildDoc(e, wikidataIssuerWords(), prov, ids, c.Langs)
	return b, rep, err == nil, err
}
