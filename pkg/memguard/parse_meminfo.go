//ff:func feature=memguard type=loader control=iteration dimension=1
//ff:what parseMeminfo — MemTotal and MemAvailable in bytes from /proc/meminfo text
package memguard

import (
	"errors"
	"strconv"
	"strings"
)

// parseMeminfo reads the two lines the guard needs. Either one missing is an
// error: the guard must not guess.
func parseMeminfo(b []byte) (total, avail int64, err error) {
	avail = -1
	for _, line := range strings.Split(string(b), "\n") {
		f := strings.Fields(line)
		if len(f) < 2 {
			continue
		}
		v, _ := strconv.ParseInt(f[1], 10, 64)
		if f[0] == "MemTotal:" {
			total = v << 10
		}
		if f[0] == "MemAvailable:" {
			avail = v << 10
		}
	}
	if total <= 0 || avail < 0 {
		return 0, 0, errors.New("memguard: no MemTotal or MemAvailable in meminfo")
	}
	return total, avail, nil
}
