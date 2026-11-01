//ff:func feature=memguard type=helper control=sequence
//ff:what slackOf — one limit's bytes above its caution, danger and recovery floors, given as shares of its size
package memguard

// slackOf measures free bytes of size against the three floors.
func slackOf(name string, free, size int64, caution, danger, recovery float64) pressure {
	over := func(share float64) int64 { return free - int64(share*float64(size)) }
	return pressure{Name: name, Free: free, Size: size, Caution: over(caution), Danger: over(danger), Recover: over(recovery)}
}
