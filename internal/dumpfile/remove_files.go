//ff:func feature=dumpfile type=writer control=iteration dimension=1
//ff:what removeFiles — remove files, ignoring ones already gone
package dumpfile

import "os"

// removeFiles removes each path; errors (a file already removed by an
// interrupted earlier attempt) are ignored: the state no longer names them.
func removeFiles(paths []string) {
	for _, p := range paths {
		os.Remove(p)
	}
}
