//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what dumpCursor.advance — move to the next entity of the sorted dump, passing over repeated lines of the same ID
package main

import (
	"fmt"
	"os"

	"github.com/geul-org/wikigeul/pkg/pipeline"
)

// advance moves cur to the next record of the merge as a dump job. A record
// with the ID of the one before (the dump repeats an entity; its first line
// comes first) is passed over and told on stderr: a document key can be put
// once, and the loaded database holds one row per ID. At the end or on an
// error ok is false.
func (d *dumpCursor) advance() {
	prev, had := d.cur, d.ok
	d.ok = false
	for {
		r, ok, err := d.m.Next()
		if err != nil || !ok {
			d.err = err
			return
		}
		if had && r.Kind == prev.Kind && r.Num == prev.Num {
			fmt.Fprintf(os.Stderr, "wikigeul: %c%d again at dump line %d, left out (the first line is kept)\n", r.Kind, r.Num, r.Line)
			continue
		}
		d.cur, d.ok = pipeline.PackJob{Kind: r.Kind, Num: r.Num, Line: string(r.Data)}, true
		return
	}
}
