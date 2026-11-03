//ff:type feature=port type=model
//ff:what Languages — the GEUL Language Registry as read from its spec document: lowercase BCP 47 tag to LANGUAGE-SHORT-CODE
package port

import "github.com/geul-org/gogeul/pkg/core"

// Languages is the GEUL Language Registry (value/text/LANGUAGE-SHORT-CODE):
// each BCP 47 tag the LANGUAGE-SHORT-CODE codec spec document
// (codec 0x00000000/codecs/0x000007a0) numbers, in lower case (tags compare without
// case), to its 1-WORD code. It is read from one version of that document
// (pipeline.LoadLanguages), never written into this code. nil is no
// registry: every language is then a LANGUAGE-BCP47-CODE.
type Languages map[string]core.Word
