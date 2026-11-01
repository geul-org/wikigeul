//ff:func feature=source type=helper control=sequence
//ff:what TestJsonbKeys — keys shorter first, then by bytes
package source

import (
	"slices"
	"testing"
)

func TestJsonbKeys(t *testing.T) {
	got := jsonbKeys(map[string]any{"precision": 1, "time": 1, "after": 1, "before": 1, "B": 1, "a": 1, "timezone": 1})
	if want := []string{"B", "a", "time", "after", "before", "timezone", "precision"}; !slices.Equal(got, want) {
		t.Fatalf("%q", got)
	}
}
