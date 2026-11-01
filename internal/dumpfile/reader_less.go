//ff:func feature=dumpfile type=helper control=sequence
//ff:what readerLess — whether one run's front record comes before another's
package dumpfile

// readerLess orders runs by their front records (compareRecords).
func readerLess(a, b *runReader) bool {
	return compareRecords(a.cur, b.cur) < 0
}
