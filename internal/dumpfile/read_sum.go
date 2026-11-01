//ff:func feature=dumpfile type=loader control=iteration dimension=1
//ff:what readSum — the hash a checksum file (sha256sum format) gives for a file name; empty when absent
package dumpfile

import (
	"os"
	"path/filepath"
	"strings"
)

// readSum reads a checksum file in the "<hash>  <name>" format (sha256sum)
// and returns the hash of the line whose name ends with name, as wikipg
// reads <dump>.sha256. A missing file or line is "".
func readSum(path, name string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(b), "\n") {
		f := strings.Fields(line)
		if len(f) == 2 && filepath.Base(strings.TrimPrefix(f[1], "*")) == name {
			return f[0]
		}
	}
	return ""
}
