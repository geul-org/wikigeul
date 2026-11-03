//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestTextKindTagSpace — over the 288 Edge codes only LABEL 0x0A4L, ALIAS 0x0A5L and DESCRIPTION 0x0A6L name a text Edge
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestTextKindTagSpace(t *testing.T) {
	names := map[uint16]string{0x0A4: "LABEL", 0x0A5: "ALIAS", 0x0A6: "DESCRIPTION"} // §3.5 Edge 번호
	for _, c := range tagSpaceCodes() {
		if c.Kind != core.KindEdge {
			continue
		}
		want, ok := names[c.Code]
		if !ok {
			want = "text"
		}
		if got := textKind(c.Codec); got != want {
			t.Errorf("0x%03XL: %q want %q", c.Code, got, want)
		}
	}
}
