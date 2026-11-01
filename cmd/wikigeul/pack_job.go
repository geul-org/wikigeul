//ff:type feature=cli type=model
//ff:what packJob — one entity to convert in pack mode, numbered in output order; from the database or, read directly from a dump, with its dump line
package main

// packJob is entity Kind+Num (Kind 'P' or 'Q') at position Seq of the output
// order (document key order: every P before every Q, then by number). Dump
// says the entity comes from a dump read directly (-dump, loadJob): Line is
// its entity JSON, empty when an -ids entity is not in the dump. Without
// Dump it is read from the loaded database.
type packJob struct {
	Seq  int64
	Kind byte
	Num  uint64
	Dump bool
	Line string
}
