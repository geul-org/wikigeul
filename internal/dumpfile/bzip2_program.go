//ff:func feature=dumpfile type=helper control=sequence
//ff:what bzip2Program — lbzip2 when it is on PATH (parallel on a single-stream dump), else bzip2
package dumpfile

import "os/exec"

// bzip2Program is the program that decompresses a .bz2 dump: lbzip2 when it
// is on PATH, else bzip2 (as wikipgeul chooses).
func bzip2Program() string {
	if _, err := exec.LookPath("lbzip2"); err == nil {
		return "lbzip2"
	}
	return "bzip2"
}
