//ff:func feature=memguard type=loader control=sequence
//ff:what readRSS — this process's resident memory in bytes from /proc/self/statm
package memguard

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// readRSS reads statm's second field (pages).
func readRSS(path string) (int64, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	f := strings.Fields(string(b))
	if len(f) < 2 {
		return 0, errors.New("memguard: statm too short")
	}
	pages, err := strconv.ParseInt(f[1], 10, 64)
	if err != nil {
		return 0, err
	}
	return pages * int64(os.Getpagesize()), nil
}
