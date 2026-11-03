//ff:func feature=port type=encoder control=sequence
//ff:what TestSourceEdge — Edge Id, Target LOCAL and Origin in order, codec 0x20
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSourceEdge(t *testing.T) {
	q, _ := entityRef("Q36578")
	r, id := sourceEdge(freshIDs(9), 2, q)
	if id != 9 || r.Codec != edgeSource || !slices.Equal(r.Payload, []core.Word{0x0051, 9, 0x0051, 2, 0x0141, 36578}) {
		t.Fatalf("%d %04x", id, r.Payload)
	}
}
