//ff:type feature=cli type=model
//ff:what dumpCursor — the sorted dump entities as jobs, one at a time: the current one, whether there is one, the read error
package main

import "github.com/geul-org/wikigeul/internal/dumpfile"

// dumpCursor walks the merged runs of a sorted dump (dumpfile.Merger) as
// dump jobs: cur is the current entity while ok; err is the error that ended
// the walk early. An ID the dump repeats gives one job, its first line
// (advance).
type dumpCursor struct {
	m   *dumpfile.Merger
	cur packJob
	ok  bool
	err error
}
