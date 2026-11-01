//ff:type feature=dumpfile type=writer
//ff:what runWriter — a run file being written: the file, its zstd encoder and a record buffer
package dumpfile

import (
	"os"

	"github.com/klauspost/compress/zstd"
)

// runWriter writes records (appendRecord) through a zstd encoder into a run
// file. A run is written under a temporary name and renamed when complete
// (writeRun, mergeRuns), so a run named in the sort state is always whole.
type runWriter struct {
	f   *os.File
	z   *zstd.Encoder
	buf []byte
}
