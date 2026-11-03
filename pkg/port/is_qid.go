//ff:func feature=port type=helper control=sequence
//ff:what isQID — whether a record is a WIKIDATA-QID Reference of any Quantifier (5Q scheme 4, core.RefScheme)
package port

import "github.com/geul-org/gogeul/pkg/core"

// isQID reports whether r is a WIKIDATA-QID Reference, whatever its
// Quantifier: Kind Reference and 5Q scheme 4 (0x014L~0x018L, core.RefScheme). A slot
// marker, pointer or ISSUER code has no scheme and is false.
func isQID(r core.Record) bool {
	s, ok := core.RefScheme(r.Codec)
	return r.Kind == core.KindReference && ok && s == core.SchemeWikidataQID
}
