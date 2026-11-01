//ff:func feature=source type=helper control=sequence
//ff:what TestWriteJSONBString — escape_json: named escapes, other control bytes as \u00xx, slash, DEL and UTF-8 as they are
package source

import (
	"strings"
	"testing"
)

func TestWriteJSONBString(t *testing.T) {
	var b strings.Builder
	writeJSONBString(&b, "a\"\\\b\f\n\r\t\x1f/\x7f한")
	if got := b.String(); got != `"a\"\\\b\f\n\r\t\u001f/`+"\x7f"+`한"` {
		t.Fatalf("got %s", got)
	}
}
