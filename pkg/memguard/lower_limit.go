//ff:func feature=memguard type=helper control=sequence
//ff:what lowerLimit — the lower of two limits, 0 meaning no limit
package memguard

// lowerLimit is the tighter of a and b.
func lowerLimit(a, b int64) int64 {
	if a == 0 || (b != 0 && b < a) {
		return b
	}
	return a
}
