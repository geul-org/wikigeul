//ff:func feature=source type=loader control=iteration dimension=1
//ff:what TestOrderedSnaks — the order list first, unlisted properties sorted after, each property's snaks in array order
package source

import "testing"

func TestOrderedSnaks(t *testing.T) {
	m := map[string][]dumpSnak{"P2": {{Property: "P2", Hash: "a"}, {Property: "P2", Hash: "b"}}, "P1": {{Property: "P1"}}, "P9": {{Property: "P9"}}, "P10": {{Property: "P10"}}}
	got := orderedSnaks([]string{"P9", "P2"}, m)
	want := []string{"P9", "P2a", "P2b", "P1", "P10"}
	if len(got) != len(want) {
		t.Fatalf("got %d snaks", len(got))
	}
	for i, s := range got {
		if s.Property+s.Hash != want[i] {
			t.Errorf("%d: got %s%s, want %s", i, s.Property, s.Hash, want[i])
		}
	}
}
