//ff:func feature=memguard type=loader control=iteration dimension=1
//ff:what findCgroup — the tightest memory limit from this process's cgroup up to the root, v2 or v1; none when there is no limit
package memguard

import (
	"os"
	"path/filepath"
)

// findCgroup reads the process's cgroup file at self and walks its directory
// under root up to the root: a child can never use more than a parent allows,
// so the lowest limit counts, measured in the cgroup that sets it.
func findCgroup(self, root string) cgroup {
	b, err := os.ReadFile(self)
	if err != nil {
		return cgroup{}
	}
	rel, v2 := cgroupPath(b)
	base := root
	if !v2 {
		base = filepath.Join(root, "memory")
	}
	var best cgroup
	for dir := filepath.Join(base, rel); ; dir = filepath.Dir(dir) {
		if c := cgroupAt(dir, v2); c.Limit > 0 && (best.Limit == 0 || c.Limit < best.Limit) {
			best = c
		}
		if dir == base || dir == filepath.Dir(dir) {
			return best
		}
	}
}
