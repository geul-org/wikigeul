//ff:func feature=dumpfile type=helper control=iteration dimension=1
//ff:what siftDown — restore the min-heap of runs below position i
package dumpfile

// siftDown moves h[i] down the binary min-heap h (readerLess) until neither
// child comes before it (leastOf).
func siftDown(h []*runReader, i int) {
	for {
		least := leastOf(h, i)
		if least == i {
			return
		}
		h[i], h[least] = h[least], h[i]
		i = least
	}
}
