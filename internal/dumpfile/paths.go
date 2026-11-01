//ff:func feature=dumpfile type=helper control=iteration dimension=1
//ff:what Paths — the run files of a sort state as paths in its directory
package dumpfile

import "path/filepath"

// Paths joins dir and each run name of a sort state (State.Runs).
func Paths(dir string, names []string) []string {
	out := make([]string, len(names))
	for i, n := range names {
		out[i] = filepath.Join(dir, n)
	}
	return out
}
