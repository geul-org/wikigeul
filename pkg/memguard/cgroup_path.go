//ff:func feature=memguard type=loader control=iteration dimension=1
//ff:what cgroupPath — this process's memory cgroup from /proc/self/cgroup text: the v2 line, or the v1 memory controller's
package memguard

import "strings"

// cgroupPath returns the path of the "0::" line (v2) or of the line whose
// controllers hold memory (v1, v2 false). A v1 memory line wins: on a hybrid
// system memory is limited there.
func cgroupPath(b []byte) (rel string, v2 bool) {
	v2 = true
	for _, line := range strings.Split(string(b), "\n") {
		f := strings.SplitN(line, ":", 3)
		if len(f) < 3 {
			continue
		}
		if f[1] == "" && v2 {
			rel = f[2]
		}
		if strings.Contains(","+f[1]+",", ",memory,") {
			return f[2], false
		}
	}
	return rel, v2
}
