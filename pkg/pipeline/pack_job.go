//ff:type feature=pipeline type=model
//ff:what PackJob — one entity to convert in pack mode, numbered in output order; read directly from a dump, with its dump line
package pipeline

// PackJob is entity Kind+Num (Kind 'P' or 'Q') at position Seq of the output
// order (document key order: every P before every Q, then by number). From a
// dump read directly Line is its entity JSON, empty when an -ids entity is
// not in the dump; the loaded database leaves it empty.
type PackJob struct {
	Seq  int64
	Kind byte
	Num  uint64
	Line string
}
