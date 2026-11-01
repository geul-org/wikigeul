//ff:func feature=dumpfile type=helper control=sequence
//ff:what sameSort — whether a saved sort state is for the same dump file and filter
package dumpfile

// sameSort compares the identity of two sort states: the dump's name, size,
// modification time and SHA-256, and the filter.
func sameSort(a, b State) bool {
	return a.Name == b.Name && a.Size == b.Size && a.ModTime == b.ModTime && a.SHA256 == b.SHA256 && a.Filter == b.Filter
}
