//ff:func feature=dumpfile type=helper control=sequence
//ff:what leastOf — which of a heap position and its two children holds the first record
package dumpfile

// leastOf returns i or the child of i in heap h whose front record comes
// first (readerLess), i when tied.
func leastOf(h []*runReader, i int) int {
	least := i
	if l := 2*i + 1; l < len(h) && readerLess(h[l], h[least]) {
		least = l
	}
	if r := 2*i + 2; r < len(h) && readerLess(h[r], h[least]) {
		least = r
	}
	return least
}
