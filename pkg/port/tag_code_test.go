//ff:type feature=port type=model
//ff:what tagCode — one code of the whole TAG space (TAG ID 0x000L~0x1FFL) with its Kind and Record Codec, read by the §3.15 section table
package port

import "github.com/geul-org/gogeul/pkg/core"

// tagCode is one code of the 512 of `000K KKKC CCCC LLLL` (migration1 §1.1
// 109, §3.15): Code is the TAG ID w>>4, Kind comes from the section KKKK
// (Code>>5: 0~2 Reference, 3·4 Value, 5~13 Edge, 14·15 Meta) and Codec is Code
// less the start of the Kind's zone (0x000, 0x060, 0x0A0, 0x1C0). The values
// are the grammar's arithmetic, not the library's constants, so tables built
// on them check the constants.
type tagCode struct {
	Code  uint16
	Kind  core.Kind
	Codec uint16
}
