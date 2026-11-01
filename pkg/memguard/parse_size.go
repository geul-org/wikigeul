//ff:func feature=memguard type=helper control=iteration dimension=1
//ff:what ParseSize — a byte count from text: a number with an optional K, M, G or T suffix (KiB, MiB …), powers of 1024
package memguard

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseSize reads "6GiB", "6G", "512MiB" or a plain number of bytes. "0" is 0.
func ParseSize(s string) (int64, error) {
	num := strings.TrimRight(strings.TrimSuffix(strings.TrimSuffix(strings.TrimSpace(s), "iB"), "B"), "KMGT")
	unit := strings.TrimPrefix(strings.TrimSpace(s), num)
	n, err := strconv.ParseInt(strings.TrimSpace(num), 10, 64)
	if err != nil || n < 0 {
		return 0, fmt.Errorf("memguard: size %q is not a number with K, M, G or T", s)
	}
	for i, u := range []string{"K", "M", "G", "T"} {
		if strings.HasPrefix(unit, u) {
			n <<= 10 * (i + 1)
		}
	}
	return n, nil
}
