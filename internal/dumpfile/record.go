//ff:type feature=dumpfile type=model
//ff:what Record — one kept entity line of the dump: kind letter, number, dump line number and the entity JSON
package dumpfile

// Record is one entity line kept from the dump: the ID's letter (P, Q) and
// number, the dump line number (1 is the "[" line, as wikipgeul counts) and the
// entity JSON without its separating comma (trimLine). Runs and the merge
// order records by (Kind, Num, Line): the documents' key order (every P
// before every Q, then the number), the first line first for an ID the dump
// repeats.
type Record struct {
	Kind byte
	Num  uint64
	Line int64
	Data []byte
}
