//ff:func feature=source type=loader control=sequence
//ff:what TestDumpSitelinks — sites by bytes from the map key, titles exact, badges in order or an empty list
package source

import (
	"reflect"
	"testing"
)

func TestDumpSitelinks(t *testing.T) {
	got := dumpSitelinks(map[string]dumpSitelink{"kowiki": {Site: "x", Title: "가"}, "enwiki": {Title: "A", Badges: []string{"Q2", "Q1"}}})
	want := []Sitelink{{Site: "enwiki", Title: "A", Badges: []string{"Q2", "Q1"}}, {Site: "kowiki", Title: "가", Badges: []string{}}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%+v", got)
	}
}
