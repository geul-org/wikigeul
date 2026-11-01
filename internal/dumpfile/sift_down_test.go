//ff:func feature=dumpfile type=helper control=sequence
//ff:what TestSiftDown — the root moves down until the heap is ordered again
package dumpfile

import "testing"

func TestSiftDown(t *testing.T) {
	r := func(n uint64) *runReader { return &runReader{cur: Record{Kind: 'Q', Num: n}} }
	h := []*runReader{r(9), r(2), r(3), r(4), r(5)}
	siftDown(h, 0)
	if h[0].cur.Num != 2 || h[1].cur.Num != 4 || h[3].cur.Num != 9 {
		t.Fatalf("%d %d %d", h[0].cur.Num, h[1].cur.Num, h[3].cur.Num)
	}
}
