//ff:type feature=memguard type=model
//ff:what pressure — how far a sample stands above the floors, over the limits in force: bytes to the caution, danger and recovery floors, and the tightest limit
package memguard

// pressure holds the bytes above each floor (negative: below it), each the
// lowest over the limits in force. Name, Free and Size describe the limit
// closest to its caution floor: "system", "cgroup" or "max".
type pressure struct {
	Name    string
	Free    int64
	Size    int64
	Caution int64
	Danger  int64
	Recover int64
}
