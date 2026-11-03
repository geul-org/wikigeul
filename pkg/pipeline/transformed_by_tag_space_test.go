//ff:func feature=pipeline type=helper control=iteration dimension=1
//ff:what TestTransformedByTagSpace — of the 64 Meta codes (0x1C0L~0x1FFL) only TRANSFORMED-BY 0x1D4L is read as the tool record of a HEAD
package pipeline

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestTransformedByTagSpace(t *testing.T) {
	for code := uint16(0x1C0); code < 0x200; code++ {
		c := code - 0x1C0 // Meta Record Codec (§2.1)
		head := []core.Record{core.EncodeSpecVersion(core.SpecVersion), {Kind: core.KindMeta, Codec: c, Payload: []core.Word{0x0051, 7}}}
		b, err := core.EncodeDocument(head, nil)
		if err != nil {
			continue // a code the encoder refuses is never in a HEAD
		}
		got := transformedBy(b)
		read := got != "" && got != core.FormatID(nil)
		if want := code == 0x1D4; read != want { // §3.5 TRANSFORMED-BY 0x1D4L
			t.Errorf("0x%03XL: read %v want %v", code, read, want)
		}
		if code == 0x1D4 && got != "0x1d4200510007" {
			t.Errorf("TRANSFORMED-BY: %q", got)
		}
	}
}
