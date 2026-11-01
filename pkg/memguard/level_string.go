//ff:func feature=memguard type=helper control=selection
//ff:what Level.String — the level's name in log lines
package memguard

// String names the level.
func (l Level) String() string {
	switch l {
	case Normal:
		return "normal"
	case Caution:
		return "caution"
	default:
		return "danger"
	}
}
