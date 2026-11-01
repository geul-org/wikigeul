//ff:func feature=source type=helper control=iteration dimension=1
//ff:what TestJsonbText — values print as value #>> '{}' of jsonb: strings unquoted, null empty, objects with keys by length then bytes
package source

import "testing"

func TestJsonbText(t *testing.T) {
	for _, c := range []struct{ in, want string }{
		{`"113230702"`, "113230702"},
		{`"a\"b\/c"`, `a"b/c`},
		{`null`, ""},
		{``, ""},
		{`{"time":"+1952-03-11T00:00:00Z","timezone":0,"before":0,"after":0,"precision":11,"calendarmodel":"http://www.wikidata.org/entity/Q1985727"}`,
			`{"time": "+1952-03-11T00:00:00Z", "after": 0, "before": 0, "timezone": 0, "precision": 11, "calendarmodel": "http://www.wikidata.org/entity/Q1985727"}`},
		{`{"text":"Douglas Adams","language":"en"}`, `{"text": "Douglas Adams", "language": "en"}`},
		{`{"latitude":52.5,"longitude":13.4,"altitude":null,"precision":1.0e-5,"globe":"http://www.wikidata.org/entity/Q2"}`,
			`{"globe": "http://www.wikidata.org/entity/Q2", "altitude": null, "latitude": 52.5, "longitude": 13.4, "precision": 0.000010}`},
		{`{"b":1,"a":[1,{"y":true,"x":"\t"}],"a":2}`, `{"a": 2, "b": 1}`},
		{`[1, {"y":true,"x":"\u0001é"}]`, "[1, {\"x\": \"\\u0001é\", \"y\": true}]"},
	} {
		got, err := jsonbText([]byte(c.in))
		if err != nil || got != c.want {
			t.Errorf("%s: got %s (%v), want %s", c.in, got, err, c.want)
		}
	}
	if _, err := jsonbText([]byte(`{"a":`)); err == nil {
		t.Fatal("not JSON: want an error")
	}
}
