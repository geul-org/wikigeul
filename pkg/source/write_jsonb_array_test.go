//ff:func feature=source type=helper control=sequence
//ff:what TestWriteJSONBArray — elements in order with ", " between; empty is []
package source

import (
	"strings"
	"testing"
)

func TestWriteJSONBArray(t *testing.T) {
	var b strings.Builder
	writeJSONBArray(&b, []any{"a", nil, []any{true}})
	writeJSONBArray(&b, nil)
	if b.String() != `["a", null, [true]][]` {
		t.Fatalf("got %s", b.String())
	}
}
