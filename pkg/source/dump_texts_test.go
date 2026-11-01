//ff:func feature=source type=loader control=sequence
//ff:what TestDumpTexts — one text per map key, keys in byte order, the key as language, NUL kept
package source

import (
	"reflect"
	"testing"
)

func TestDumpTexts(t *testing.T) {
	got := dumpTexts(map[string]dumpTerm{"mul": {Language: "x", Value: "M"}, "en": {Value: "a\x00b"}, "de-formal": {Value: "D"}})
	if !reflect.DeepEqual(got, []Text{{"de-formal", "D"}, {"en", "a\x00b"}, {"mul", "M"}}) {
		t.Fatalf("%+v", got)
	}
}
