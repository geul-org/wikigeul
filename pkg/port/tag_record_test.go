//ff:func feature=port type=helper control=sequence
//ff:what tagRecord — a Record of a TAG space code with one payload word
package port

import "github.com/geul-org/gogeul/pkg/core"

// tagRecord is a record of c with the payload word 5, enough for judgments
// that read only the Kind and Codec.
func tagRecord(c tagCode) core.Record {
	return core.Record{Kind: c.Kind, Codec: c.Codec, Payload: []core.Word{5}}
}
