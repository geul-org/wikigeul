//ff:func feature=source type=loader control=iteration dimension=1
//ff:what TestDumpSnakValue — no datavalue, \u0000 values as wikipg keeps them, entity IDs, jsonb text, and a value that is not JSON
package source

import (
	"encoding/json"
	"testing"
)

func TestDumpSnakValue(t *testing.T) {
	if typ, v, err := dumpSnakValue(nil); typ != "" || v != "" || err != nil {
		t.Fatal("nil")
	}
	for _, c := range []struct{ typ, raw, want string }{
		{"string", `"x\u0000"`, "x\x00"},
		{"monolingualtext", `{"text":"a\u0000","language":"en"}`, `{"text":"a\u0000","language":"en"}`},
		{"wikibase-entityid", `{"entity-type":"item","id":"Q5"}`, "Q5"},
		{"wikibase-entityid", `{"entity-type":"item","numeric-id":5}`, `{"numeric-id": 5, "entity-type": "item"}`},
		{"string", `"a"`, "a"},
		{"quantity", `{"amount":"+1","unit":"1"}`, `{"unit": "1", "amount": "+1"}`},
	} {
		typ, v, err := dumpSnakValue(&dumpDataValue{Type: c.typ, Value: json.RawMessage(c.raw)})
		if typ != c.typ || v != c.want || err != nil {
			t.Errorf("%s: got %s %q %v", c.raw, typ, v, err)
		}
	}
	if _, _, err := dumpSnakValue(&dumpDataValue{Type: "string", Value: json.RawMessage(`{`)}); err == nil {
		t.Fatal("not JSON: want an error")
	}
}
