//ff:func feature=cli type=helper control=sequence
//ff:what TestFirstQ — nil for no list, the first n Q entities of a list, an empty non-nil slice for a list without Q
package main

import "testing"

func TestFirstQ(t *testing.T) {
	if firstQ(nil, 2) != nil {
		t.Fatal("no list")
	}
	got := firstQ([]packJob{{Kind: 'P', Num: 1}, {Kind: 'Q', Num: 1}, {Kind: 'Q', Num: 2}, {Kind: 'Q', Num: 3}}, 2)
	if len(got) != 2 || got[1].Num != 2 {
		t.Fatalf("%+v", got)
	}
	if got := firstQ([]packJob{{Kind: 'P', Num: 1}}, 2); got == nil || len(got) != 0 {
		t.Fatalf("%+v", got)
	}
}
