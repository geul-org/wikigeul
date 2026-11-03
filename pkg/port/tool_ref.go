//ff:func feature=port type=encoder control=sequence
//ff:what toolRef — the EXTERNAL-ISSUER-2WORD Reference to LOCAL 2 of geul.org term definition document n
package port

import "github.com/geul-org/gogeul/pkg/core"

// toolRef points at the tool that geul.org document 0x0001000000000000+n
// defines (ISSUER 용어 정의 문서, one tool per document: LOCAL 2 is the tool) as
// EXTERNAL-ISSUER-2WORD × Identified (0x02DL): the GEUL Issuer 0x00000000 in
// 2 words, the 4-word document ID and LOCAL 2.
func toolRef(n uint16) core.Record {
	return core.Record{Kind: core.KindReference, Codec: core.CodecExternalIssuer2Word, Payload: []core.Word{0, 0, 0x0001, 0, 0, core.Word(n), 2}}
}
