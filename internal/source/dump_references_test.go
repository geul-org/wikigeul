//ff:func feature=source type=loader control=sequence
//ff:what TestDumpReferences — references in order with their ordinal and hash, snaks in snaks-order; empty references left out; a bad value is an error
package source

import (
	"encoding/json"
	"testing"
)

func TestDumpReferences(t *testing.T) {
	str := func(s string) *dumpDataValue {
		return &dumpDataValue{Type: "string", Value: json.RawMessage(`"` + s + `"`)}
	}
	rs, err := dumpReferences(4, []dumpReference{
		{Hash: "empty"},
		{Hash: "h", Snaks: map[string][]dumpSnak{"P854": {{Property: "P854", DataValue: str("u")}}, "P813": {{Property: "P813", DataValue: str("d")}}}, SnaksOrder: []string{"P854", "P813"}},
	})
	if err != nil || len(rs) != 1 || rs[0].Ordinal != 1 || rs[0].Hash != "h" || rs[0].StatementID != 4 || rs[0].Snaks[0].Value != "u" || rs[0].Snaks[1].Value != "d" {
		t.Fatalf("%+v %v", rs, err)
	}
	bad := []dumpReference{{Snaks: map[string][]dumpSnak{"P1": {{DataValue: &dumpDataValue{Value: json.RawMessage(`{`)}}}}}}
	if _, err := dumpReferences(4, bad); err == nil {
		t.Fatal("want an error")
	}
}
