//ff:func feature=source type=loader control=sequence
//ff:what TestDecodeDumpEntity — a dump line decodes to the Entity the PG queries give: orders, value forms, references, sitelinks, NUL kept
package source

import (
	"reflect"
	"testing"
)

const testDumpLine = `{"type":"item","id":"Q42","lastrevid":2557857485,"modified":"2026-10-17T04:21:25Z",` +
	`"labels":{"mul":{"language":"mul","value":"Douglas Adams"},"en":{"language":"en","value":"Douglas Adams"},"ko":{"language":"ko","value":"더글러스 애덤스"},"x":{"language":"x","value":"a\u0000b"}},` +
	`"descriptions":{"en":{"language":"en","value":"English writer"}},` +
	`"aliases":{"en":[{"language":"en","value":"DNA"},{"language":"en","value":"Douglas Noël Adams"}]},` +
	`"claims":{"P569":[{"mainsnak":{"snaktype":"value","property":"P569","datavalue":{"value":{"time":"+1952-03-11T00:00:00Z","timezone":0,"before":0,"after":0,"precision":11,"calendarmodel":"http://www.wikidata.org/entity/Q1985727"},"type":"time"},"datatype":"time"},"type":"statement","id":"Q42$b","rank":"normal"}],` +
	`"P31":[{"mainsnak":{"snaktype":"value","property":"P31","datavalue":{"value":{"entity-type":"item","numeric-id":5,"id":"Q5"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement",` +
	`"qualifiers":{"P642":[{"snaktype":"value","property":"P642","datavalue":{"value":{"entity-type":"item","numeric-id":1,"id":"Q1"},"type":"wikibase-entityid"},"datatype":"wikibase-item"}],"P582":[{"snaktype":"novalue","property":"P582","datatype":"time"}]},` +
	`"qualifiers-order":["P642","P582"],"id":"Q42$a","rank":"preferred",` +
	`"references":[{"hash":"h1","snaks":{"P854":[{"snaktype":"value","property":"P854","datavalue":{"value":"https://a","type":"string"},"datatype":"url"}],"P248":[{"snaktype":"value","property":"P248","datavalue":{"value":{"entity-type":"item","numeric-id":9,"id":"Q9"},"type":"wikibase-entityid"},"datatype":"wikibase-item"}]},"snaks-order":["P854","P248"]},{"hash":"h2","snaks":{},"snaks-order":[]}]}],` +
	`"P214":[{"mainsnak":{"snaktype":"value","property":"P214","datavalue":{"value":"x\u0000","type":"string"},"datatype":"external-id"},"type":"statement","id":"Q42$c","rank":"normal"}]},` +
	`"sitelinks":{"kowiki":{"site":"kowiki","title":"더글러스 애덤스","badges":[]},"enwiki":{"site":"enwiki","title":"Douglas Adams","badges":["Q17437796"]}}}`

func TestDecodeDumpEntity(t *testing.T) {
	e, err := DecodeDumpEntity([]byte(testDumpLine))
	if err != nil {
		t.Fatal(err)
	}
	if e.ID != "Q42" || e.LastRevID != 2557857485 || e.Modified != 1792210885 || e.Datatype != "" {
		t.Fatalf("head %+v", e)
	}
	if !reflect.DeepEqual(e.Labels, []Text{{"en", "Douglas Adams"}, {"ko", "더글러스 애덤스"}, {"mul", "Douglas Adams"}, {"x", "a\x00b"}}) {
		t.Fatalf("labels %+v", e.Labels)
	}
	if !reflect.DeepEqual(e.Aliases, []Text{{"en", "DNA"}, {"en", "Douglas Noël Adams"}}) {
		t.Fatalf("aliases %+v", e.Aliases)
	}
	if len(e.Claims) != 3 || e.Claims[0].Property != "P214" || e.Claims[1].Property != "P31" || e.Claims[2].Property != "P569" {
		t.Fatalf("claim order %+v", e.Claims)
	}
	if e.Claims[0].Value != "x\x00" || e.Claims[1].Value != "Q5" || e.Claims[1].Rank != "preferred" || e.Claims[1].Datatype != "wikibase-item" {
		t.Fatalf("values %+v", e.Claims)
	}
	if e.Claims[2].Value != `{"time": "+1952-03-11T00:00:00Z", "after": 0, "before": 0, "timezone": 0, "precision": 11, "calendarmodel": "http://www.wikidata.org/entity/Q1985727"}` {
		t.Fatalf("time %q", e.Claims[2].Value)
	}
	qs := e.Claims[1].Qualifiers
	if len(qs) != 2 || qs[0].Property != "P642" || qs[0].Value != "Q1" || qs[1].SnakType != "novalue" || qs[1].Value != "" || qs[0].StatementID != 1 {
		t.Fatalf("qualifiers %+v", qs)
	}
	rs := e.Claims[1].References
	if len(rs) != 1 || rs[0].Hash != "h1" || rs[0].Ordinal != 0 || len(rs[0].Snaks) != 2 || rs[0].Snaks[0].Value != "https://a" || rs[0].Snaks[1].Value != "Q9" {
		t.Fatalf("references %+v", rs)
	}
	if len(e.Sitelinks) != 2 || e.Sitelinks[0].Site != "enwiki" || e.Sitelinks[0].Badges[0] != "Q17437796" || e.Sitelinks[1].Badges == nil {
		t.Fatalf("sitelinks %+v", e.Sitelinks)
	}
	if _, err := DecodeDumpEntity([]byte(`{"type":"item","labels":[]}`)); err == nil {
		t.Fatal("labels as an array: want an error")
	}
	if _, err := DecodeDumpEntity([]byte(`{"type":"item"}`)); err == nil {
		t.Fatal("no id: want an error")
	}
	if _, err := DecodeDumpEntity([]byte(`{"id":"Q1","modified":"yesterday"}`)); err == nil {
		t.Fatal("bad modified: want an error")
	}
}
