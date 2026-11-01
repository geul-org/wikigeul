//ff:type feature=memguard type=model
//ff:what cgroup — the tightest memory cgroup of this process: its limit and the files that say what it uses
package memguard

// cgroup is the cgroup with the lowest limit among this process's and its
// parents'; Limit 0 is none. Current holds its memory use, Stat the
// statistics whose line Key counts the inactive file cache.
type cgroup struct {
	Limit   int64
	Current string
	Stat    string
	Key     string
}
