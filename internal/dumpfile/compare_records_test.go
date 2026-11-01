//ff:func feature=dumpfile type=helper control=sequence
//ff:what TestCompareRecords — kind first (P before Q), then number, then dump line
package dumpfile

import "testing"

func TestCompareRecords(t *testing.T) {
	p, q1, q1b, q2 := Record{Kind: 'P', Num: 9}, Record{Kind: 'Q', Num: 1, Line: 5}, Record{Kind: 'Q', Num: 1, Line: 6}, Record{Kind: 'Q', Num: 2}
	if compareRecords(p, q1) >= 0 || compareRecords(q1, q2) >= 0 || compareRecords(q1, q1b) >= 0 || compareRecords(q2, q1) <= 0 || compareRecords(q1, q1) != 0 {
		t.Fatal("order")
	}
}
