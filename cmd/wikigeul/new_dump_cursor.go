//ff:func feature=cli type=helper control=sequence
//ff:what newDumpCursor — a cursor at the first entity of a sorted dump
package main

import "github.com/geul-org/wikigeul/internal/dumpfile"

// newDumpCursor starts a dumpCursor over m at its first entity.
func newDumpCursor(m *dumpfile.Merger) *dumpCursor {
	d := &dumpCursor{m: m}
	d.advance()
	return d
}
