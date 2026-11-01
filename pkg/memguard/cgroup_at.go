//ff:func feature=memguard type=loader control=selection
//ff:what cgroupAt — the memory limit of one cgroup directory with the files that measure it: v2 the lower of max and high, v1 the limit
package memguard

import "path/filepath"

// cgroupAt reads dir's limit; "max", no file or v1's huge number is Limit 0.
func cgroupAt(dir string, v2 bool) cgroup {
	switch {
	case v2:
		return cgroup{Limit: lowerLimit(readLimit(filepath.Join(dir, "memory.max")), readLimit(filepath.Join(dir, "memory.high"))),
			Current: filepath.Join(dir, "memory.current"), Stat: filepath.Join(dir, "memory.stat"), Key: "inactive_file"}
	default:
		return cgroup{Limit: readLimit(filepath.Join(dir, "memory.limit_in_bytes")),
			Current: filepath.Join(dir, "memory.usage_in_bytes"), Stat: filepath.Join(dir, "memory.stat"), Key: "total_inactive_file"}
	}
}
