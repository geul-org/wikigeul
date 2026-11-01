//ff:func feature=source type=loader control=sequence
//ff:what dumpModified — a dump line's modified time as POSIX seconds, 0 when absent
package source

import (
	"fmt"
	"time"
)

// dumpModified parses the modified field of a dump line (RFC 3339, e.g.
// 2026-10-18T10:23:45Z) as the PG query gives entity.modified:
// floor(extract(epoch …)), 0 when the field is empty (wikipgeul stores NULL).
// A time PostgreSQL would also reject is an error.
func dumpModified(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return 0, fmt.Errorf("wikigeul: dump modified %q: %w", s, err)
	}
	return t.Unix(), nil
}
