//ff:func feature=memguard type=helper control=sequence
//ff:what fixtureAt — the time sec seconds into a test
package memguard

import "time"

// fixtureAt is a fixed start plus sec seconds.
func fixtureAt(sec int) time.Time {
	return time.Unix(1_800_000_000+int64(sec), 0)
}
