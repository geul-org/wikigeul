//ff:func feature=port type=helper control=iteration dimension=1
//ff:what tagSpaceCodes — all 512 codes of the TAG space (0x000L~0x1FFL) with their Kind and Record Codec
package port

import "github.com/geul-org/gogeul/pkg/core"

// tagSpaceCodes lists every code 0x000~0x1FF in order (tagCode).
func tagSpaceCodes() []tagCode {
	zones := [8]struct {
		kind core.Kind
		base uint16
	}{{core.KindReference, 0x000}, {core.KindValue, 0x040}, {core.KindEdge, 0x080}, {core.KindEdge, 0x080}, {core.KindEdge, 0x080}, {core.KindEdge, 0x080}, {core.KindEdge, 0x080}, {core.KindMeta, 0x1C0}}
	out := make([]tagCode, 0, 512)
	for code := uint16(0); code < 0x200; code++ {
		z := zones[code>>6]
		out = append(out, tagCode{Code: code, Kind: z.kind, Codec: code - z.base})
	}
	return out
}
