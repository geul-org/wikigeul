//ff:func feature=source type=loader control=sequence
//ff:what TestDumpAliases — languages by bytes, each language's aliases in array order; none gives none
package source

import (
	"reflect"
	"testing"
)

func TestDumpAliases(t *testing.T) {
	got := dumpAliases(map[string][]dumpTerm{"ko": {{Value: "가"}}, "en": {{Value: "b"}, {Value: "a"}}})
	if !reflect.DeepEqual(got, []Text{{"en", "b"}, {"en", "a"}, {"ko", "가"}}) {
		t.Fatalf("%+v", got)
	}
	if dumpAliases(nil) != nil {
		t.Fatal("none")
	}
}
