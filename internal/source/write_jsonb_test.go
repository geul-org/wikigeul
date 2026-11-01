//ff:func feature=source type=helper control=iteration dimension=1
//ff:what TestWriteJSONB — every decoded JSON kind printed as jsonb prints it
package source

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestWriteJSONB(t *testing.T) {
	for _, c := range []struct {
		v    any
		want string
	}{
		{nil, "null"}, {true, "true"}, {false, "false"}, {json.Number("1.0e-2"), "0.010"}, {"a\nb", `"a\nb"`},
		{[]any{json.Number("1"), "x"}, `[1, "x"]`}, {map[string]any{"bb": nil, "a": []any{}}, `{"a": [], "bb": null}`}, {map[string]any{}, "{}"},
	} {
		var b strings.Builder
		writeJSONB(&b, c.v)
		if b.String() != c.want {
			t.Errorf("%v: got %s, want %s", c.v, b.String(), c.want)
		}
	}
}
