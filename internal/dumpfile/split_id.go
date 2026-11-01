//ff:func feature=dumpfile type=loader control=sequence
//ff:what SplitID — an entity ID's letter and number (Q42 → Q, 42) as wikipgeul splits it; false for anything else
package dumpfile

import "strconv"

// SplitID splits an entity ID such as Q42 or P31 into its letter and
// positive number exactly as wikipgeul does (dump.SplitID): a leading zero, a
// sign, a non-number or a number beyond int64 is false. wikipgeul loads no
// entity for such an ID (load_error), so the dump path skips it too.
func SplitID(id string) (byte, uint64, bool) {
	if len(id) < 2 || id[1] < '1' || id[1] > '9' {
		return 0, 0, false
	}
	n, err := strconv.ParseInt(id[1:], 10, 64)
	if err != nil {
		return 0, 0, false
	}
	return id[0], uint64(n), true
}
