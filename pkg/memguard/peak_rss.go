//ff:func feature=memguard type=loader control=iteration dimension=1
//ff:what peakRSS — the most resident memory this process has had, in bytes, from VmHWM in /proc/self/status; 0 when unknown
package memguard

import (
	"os"
	"strconv"
	"strings"
)

// peakRSS reads the VmHWM line (kB) of the status file at path.
func peakRSS(path string) int64 {
	b, _ := os.ReadFile(path)
	for _, line := range strings.Split(string(b), "\n") {
		f := strings.Fields(line)
		if len(f) >= 2 && f[0] == "VmHWM:" {
			v, _ := strconv.ParseInt(f[1], 10, 64)
			return v << 10
		}
	}
	return 0
}
