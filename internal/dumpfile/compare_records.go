//ff:func feature=dumpfile type=helper control=sequence
//ff:what compareRecords — run order of two records: kind letter, number, then dump line
package dumpfile

import "cmp"

// compareRecords orders a before b by kind letter (P before Q, as their
// documents 0x50… and 0x51…), number, then dump line.
func compareRecords(a, b Record) int {
	return cmp.Or(cmp.Compare(a.Kind, b.Kind), cmp.Compare(a.Num, b.Num), cmp.Compare(a.Line, b.Line))
}
