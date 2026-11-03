//ff:func feature=port type=helper control=sequence
//ff:what isQID — whether a record is a WIKIDATA-QID Reference of any Quantifier (scheme C/5 = 4)
package port

import "github.com/geul-org/gogeul/pkg/core"

// isQID reports whether r is a WIKIDATA-QID Reference, whatever its
// Quantifier: Kind Reference and scheme C/5 = 4 (0x014L~0x018L). A slot
// marker, pointer or ISSUER code has no scheme and is false.
func isQID(r core.Record) bool {
	s, ok := core.RefScheme(r.Codec)
	return r.Kind == core.KindReference && ok && s == core.SchemeWikidataQID
}
