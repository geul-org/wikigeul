//ff:func feature=memguard type=helper control=sequence
//ff:what bump — a count of events in a row: one more, or back to zero
package memguard

// bump is n+1 when hit, else 0.
func bump(n int, hit bool) int {
	if hit {
		return n + 1
	}
	return 0
}
