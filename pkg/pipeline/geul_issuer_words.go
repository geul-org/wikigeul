//ff:func feature=pipeline type=helper control=sequence
//ff:what geulIssuerWords — the GEUL Issuer (geul.org) ID 0x00000000 as its two words
package pipeline

import "github.com/geul-org/gogeul/pkg/core"

// geulIssuerWords returns the GEUL Issuer geul.org, 0x00000000, as 2 words:
// the Issuer of the codec documents (`0x00000000/codecs/…`). A new slice each
// call.
func geulIssuerWords() []core.Word {
	return []core.Word{0, 0}
}
