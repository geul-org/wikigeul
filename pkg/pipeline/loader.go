//ff:type feature=pipeline type=model
//ff:what Loader — where a job's entity comes from: the one thing the two conversion inputs differ in
package pipeline

import "github.com/geul-org/wikigeul/pkg/source"

// Loader reads job's entity as source.Entity, the one input port.BuildDoc
// converts: from the loaded database (wikipg export-geul, one Loader per
// worker with its own connection) or from its line of a dump read directly
// (wikigeul). Both give the same Entity for the same entity (docs/dump.md),
// which is why everything after it is the same code. false, with no error,
// is an entity the input does not hold.
type Loader func(job PackJob) (source.Entity, bool, error)
