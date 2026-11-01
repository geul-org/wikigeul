//ff:func feature=memguard type=loader control=sequence
//ff:what readLimit — a cgroup limit file as bytes; 0 for "max", no file, or v1's "no limit" near 2^63
package memguard

import (
	"os"
	"strconv"
	"strings"
)

// readLimit reads one number. Anything from 2^62 up means no limit (v1).
func readLimit(path string) int64 {
	b, err := os.ReadFile(path)
	if err != nil {
		return 0
	}
	v, err := strconv.ParseInt(strings.TrimSpace(string(b)), 10, 64)
	if err != nil || v >= 1<<62 {
		return 0
	}
	return v
}
