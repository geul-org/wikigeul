//ff:func feature=port type=helper control=iteration dimension=1
//ff:what tagCodesOf — the codes of the TAG space whose Kind and Record Codec are k and codec
package port

import "github.com/geul-org/gogeul/pkg/core"

// tagCodesOf returns every code of tagSpaceCodes that is k/codec.
func tagCodesOf(k core.Kind, codec uint16) []uint16 {
	var out []uint16
	for _, c := range tagSpaceCodes() {
		if c.Kind == k && c.Codec == codec {
			out = append(out, c.Code)
		}
	}
	return out
}
