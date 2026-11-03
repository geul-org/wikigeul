//ff:func feature=pipeline type=helper control=sequence
//ff:what wikidataIssuerWords — wikidata.geul.org's Issuer ID 0x00000001 as its two words
package pipeline

import "github.com/geul-org/gogeul/pkg/core"

// wikidataIssuerWords returns WikidataIssuer as the 2-word Issuer
// (0x0000 0x0001) records and pack keys carry. A new slice each call, so no
// caller can change another's.
func wikidataIssuerWords() []core.Word {
	return []core.Word{0, 1}
}
