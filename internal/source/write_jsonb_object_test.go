//ff:func feature=source type=helper control=sequence
//ff:what TestWriteJSONBObject — keys in jsonb order with ": " and ", ", keys escaped; empty is {}
package source

import (
	"strings"
	"testing"
)

func TestWriteJSONBObject(t *testing.T) {
	var b strings.Builder
	writeJSONBObject(&b, map[string]any{"text": "x", "language": "en", "a\"": map[string]any{}})
	writeJSONBObject(&b, nil)
	if b.String() != `{"a\"": {}, "text": "x", "language": "en"}{}` {
		t.Fatalf("got %s", b.String())
	}
}
