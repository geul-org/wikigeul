//ff:func feature=cli type=command control=selection
//ff:what loadJob — the entity of a job from where it comes: the loaded database, or its line of a dump read directly
package main

import (
	"database/sql"
	"strconv"

	"github.com/geul-org/wikigeul/internal/source"
)

// loadJob reads entity job as source.Entity, the one input port.BuildDoc
// converts: from the loaded database (source.LoadEntity, each reference
// occurrence in its own snak order when c.RefOrder), or for a dump job from
// its dump line (decodeJob), which gives the same Entity for the same
// entity (docs/dump.md). A dump job with no line is an -ids entity the dump
// does not hold: false, as the database gives for it.
func loadJob(db *sql.DB, c config, job packJob) (source.Entity, bool, error) {
	switch {
	case !job.Dump:
		id := string(job.Kind) + strconv.FormatUint(job.Num, 10)
		return source.LoadEntity(db, source.Store{Schema: c.Schema, Dump: c.Dump, RefOrder: c.RefOrder}, id)
	case job.Line == "":
		return source.Entity{}, false, nil
	default:
		return decodeJob(job)
	}
}
