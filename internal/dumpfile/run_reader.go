//ff:type feature=dumpfile type=loader
//ff:what runReader — an open run file and its current (smallest unread) record
package dumpfile

import (
	"bufio"
	"os"

	"github.com/klauspost/compress/zstd"
)

// runReader reads one run file in order: cur is the record at its front,
// which a merge compares (readerLess).
type runReader struct {
	f   *os.File
	z   *zstd.Decoder
	br  *bufio.Reader
	cur Record
}
