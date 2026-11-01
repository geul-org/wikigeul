//ff:func feature=source type=loader control=iteration dimension=1
//ff:what TestDumpClaim — a statement as a Claim with its main snak, rank, qualifiers and references; a bad value anywhere is an error
package source

import (
	"encoding/json"
	"testing"
)

func TestDumpClaim(t *testing.T) {
	ok := &dumpDataValue{Type: "string", Value: json.RawMessage(`"v"`)}
	bad := &dumpDataValue{Value: json.RawMessage(`{`)}
	s := dumpStatement{Rank: "preferred", Mainsnak: dumpSnak{SnakType: "value", Property: "P1", Datatype: "string", DataValue: ok},
		Qualifiers: map[string][]dumpSnak{"P2": {{Property: "P2", DataValue: ok}}},
		References: []dumpReference{{Hash: "h", Snaks: map[string][]dumpSnak{"P3": {{Property: "P3", DataValue: ok}}}}}}
	c, err := dumpClaim(9, s)
	if err != nil || c.ID != 9 || c.Property != "P1" || c.Value != "v" || c.Rank != "preferred" || c.Type != "string" || c.Datatype != "string" || c.SnakType != "value" ||
		len(c.Qualifiers) != 1 || len(c.References) != 1 {
		t.Fatalf("%+v %v", c, err)
	}
	for _, b := range []dumpStatement{
		{Mainsnak: dumpSnak{DataValue: bad}},
		{Qualifiers: map[string][]dumpSnak{"P2": {{DataValue: bad}}}},
		{References: []dumpReference{{Snaks: map[string][]dumpSnak{"P3": {{DataValue: bad}}}}}},
	} {
		if _, err := dumpClaim(9, b); err == nil {
			t.Errorf("%+v: want an error", b)
		}
	}
}
