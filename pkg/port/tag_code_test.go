//ff:type feature=port type=model
//ff:what tagCode — one code of the whole TAG space (TAG ID 0x000L~0x1FFL) with its Kind and Record Codec, read by the §2.1 arithmetic
package port

import "github.com/geul-org/gogeul/pkg/core"

// tagCode is one code of the 512 of `000K KKCC CCCC LLLL` (migration1 §2.1):
// Code is the TAG ID w>>4, Kind comes from KKK (Code>>6: 0 Reference, 1 Value,
// 2~6 Edge, 7 Meta) and Codec is Code less the start of the Kind's zone
// (0x000, 0x040, 0x080, 0x1C0). The values are the grammar's arithmetic, not
// the library's constants, so tables built on them check the constants.
type tagCode struct {
	Code  uint16
	Kind  core.Kind
	Codec uint16
}
