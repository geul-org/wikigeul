//ff:func feature=port type=encoder control=sequence
//ff:what TestSnakMarker — somevalue is UNKNOWN-VALUE 0x0010, novalue is NO-VALUE 0x0020, a value snak and an empty type have none
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSnakMarker(t *testing.T) {
	if r, ok := snakMarker("somevalue"); !ok || !core.IsSlotMarker(r) || !slices.Equal(wordsOf(r), []core.Word{0x0010}) {
		t.Fatalf("somevalue: %+v %v", r, ok)
	}
	if r, ok := snakMarker("novalue"); !ok || !core.IsSlotMarker(r) || !slices.Equal(wordsOf(r), []core.Word{0x0020}) {
		t.Fatalf("novalue: %+v %v", r, ok)
	}
	if _, ok := snakMarker("value"); ok {
		t.Fatal("value: want false")
	}
	if _, ok := snakMarker(""); ok {
		t.Fatal("empty: want false")
	}
}
