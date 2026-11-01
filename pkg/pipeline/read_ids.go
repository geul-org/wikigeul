//ff:func feature=pipeline type=helper control=iteration dimension=1
//ff:what ReadIDs — the entities of an ID list file in document key order, without duplicates, with the file's SHA-256
package pipeline

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// ReadIDs reads path (one Q… or P… per line; blank lines and # comments
// skipped) and returns the entities sorted as their documents are keyed
// (JobAfter), duplicates removed, and the file's SHA-256 in hex, which the
// checkpoint records. Path "" is no list. Another kind or number is an error.
func ReadIDs(path string) ([]PackJob, string, error) {
	if path == "" {
		return nil, "", nil
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, "", err
	}
	var out []PackJob
	for _, l := range strings.Split(string(b), "\n") {
		l = strings.TrimSpace(l)
		if l == "" || l[0] == '#' {
			continue
		}
		n, err := strconv.ParseUint(l[1:], 10, 64)
		if err != nil || (l[0] != 'Q' && l[0] != 'P') || n == 0 {
			return nil, "", fmt.Errorf("wikigeul: -ids: %q is not a Q or P ID", l)
		}
		out = append(out, PackJob{Kind: l[0], Num: n})
	}
	slices.SortFunc(out, func(a, b PackJob) int {
		if JobAfter(a.Kind, a.Num, b.Kind, b.Num) {
			return 1
		}
		if a == b {
			return 0
		}
		return -1
	})
	sum := sha256.Sum256(b)
	return slices.Compact(out), hex.EncodeToString(sum[:]), nil
}
