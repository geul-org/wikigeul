//ff:func feature=memguard type=loader control=iteration dimension=1
//ff:what cgroup.used — the cgroup's memory use less its inactive file cache, which the kernel reclaims first
package memguard

import (
	"os"
	"strconv"
	"strings"
)

// used reads Current and takes off the Key line of Stat, as a container's
// working set is counted.
func (c cgroup) used() (int64, error) {
	b, err := os.ReadFile(c.Current)
	if err != nil {
		return 0, err
	}
	cur, err := strconv.ParseInt(strings.TrimSpace(string(b)), 10, 64)
	if err != nil {
		return 0, err
	}
	stat, err := os.ReadFile(c.Stat)
	if err != nil {
		return 0, err
	}
	for _, line := range strings.Split(string(stat), "\n") {
		f := strings.Fields(line)
		if len(f) == 2 && f[0] == c.Key {
			v, _ := strconv.ParseInt(f[1], 10, 64)
			return max(0, cur-v), nil
		}
	}
	return cur, nil
}
