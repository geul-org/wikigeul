//ff:func feature=dumpfile type=helper control=sequence
//ff:what TestLeastOf — the position itself, its left or its right child, whichever front record comes first; children past the end are ignored
package dumpfile

import "testing"

func TestLeastOf(t *testing.T) {
	r := func(n uint64) *runReader { return &runReader{cur: Record{Kind: 'Q', Num: n}} }
	if got := leastOf([]*runReader{r(1), r(2), r(3)}, 0); got != 0 {
		t.Fatalf("self: %d", got)
	}
	if got := leastOf([]*runReader{r(5), r(2), r(3)}, 0); got != 1 {
		t.Fatalf("left: %d", got)
	}
	if got := leastOf([]*runReader{r(5), r(4), r(3)}, 0); got != 2 {
		t.Fatalf("right: %d", got)
	}
	if got := leastOf([]*runReader{r(5), r(4)}, 1); got != 1 {
		t.Fatalf("leaf: %d", got)
	}
}
