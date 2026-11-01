//ff:func feature=dumpfile type=helper control=sequence
//ff:what TestReaderLess — runs compare by their front records
package dumpfile

import "testing"

func TestReaderLess(t *testing.T) {
	a, b := &runReader{cur: Record{Kind: 'P', Num: 5}}, &runReader{cur: Record{Kind: 'Q', Num: 1}}
	if !readerLess(a, b) || readerLess(b, a) || readerLess(a, a) {
		t.Fatal("order")
	}
}
