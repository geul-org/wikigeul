//ff:func feature=cli type=command control=selection
//ff:what loadJob — the entity of a job from its line of the dump; missing for a job with no line
package main

import (
	"github.com/geul-org/wikigeul/pkg/pipeline"
	"github.com/geul-org/wikigeul/pkg/source"
)

// loadJob is the dump's pipeline.Loader: job's entity from its dump line
// (decodeJob), which gives the Entity the loaded database would give for the
// same entity (docs/dump.md). A job with no line is an -ids entity the dump
// does not hold: false, as the database gives for it.
func loadJob(job pipeline.PackJob) (source.Entity, bool, error) {
	switch {
	case job.Line == "":
		return source.Entity{}, false, nil
	default:
		return decodeJob(job)
	}
}
