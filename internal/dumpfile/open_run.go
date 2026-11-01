//ff:func feature=dumpfile type=loader control=sequence
//ff:what openRun — open a run file positioned at its first record; false for an empty run
package dumpfile

import (
	"bufio"
	"os"

	"github.com/klauspost/compress/zstd"
)

// openRun opens the run at path with a single-goroutine, low-memory zstd
// decoder and reads its first record. An empty run is false (and closed).
func openRun(path string) (*runReader, bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, false, err
	}
	z, err := zstd.NewReader(f, zstd.WithDecoderConcurrency(1), zstd.WithDecoderLowmem(true))
	if err != nil {
		f.Close()
		return nil, false, err
	}
	rr := &runReader{f: f, z: z, br: bufio.NewReaderSize(z, 256<<10)}
	ok, err := rr.advance()
	if err != nil || !ok {
		rr.close()
	}
	return rr, ok, err
}
