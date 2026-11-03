//ff:func feature=port type=helper control=iteration dimension=1
//ff:what tagSpaceCodes — all 512 codes of the TAG space (0x000L~0x1FFL) with their Kind and Record Codec
package port

import "github.com/geul-org/gogeul/pkg/core"

// tagSpaceCodes lists every code 0x000~0x1FF in order (tagCode).
func tagSpaceCodes() []tagCode {
	const r, v, e, m = core.KindReference, core.KindValue, core.KindEdge, core.KindMeta
	kinds := [16]core.Kind{r, r, r, v, v, e, e, e, e, e, e, e, e, e, m, m} // migration1 §3.15, by section KKKK
	base := map[core.Kind]uint16{r: 0x000, v: 0x060, e: 0x0A0, m: 0x1C0}
	out := make([]tagCode, 0, 512)
	for code := uint16(0); code < 0x200; code++ {
		k := kinds[code>>5]
		out = append(out, tagCode{Code: code, Kind: k, Codec: code - base[k]})
	}
	return out
}
