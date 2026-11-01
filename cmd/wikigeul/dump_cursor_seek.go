//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what dumpCursor.seek — move to an entity of the sorted dump; whether the dump holds it
package main

// seek advances past the entities before (kind, num) in key order (jobAfter)
// and reports whether the cursor is then on that entity.
func (d *dumpCursor) seek(kind byte, num uint64) bool {
	for d.ok && jobAfter(kind, num, d.cur.Kind, d.cur.Num) {
		d.advance()
	}
	return d.ok && d.cur.Kind == kind && d.cur.Num == num
}
