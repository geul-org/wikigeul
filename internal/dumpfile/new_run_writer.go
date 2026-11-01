//ff:func feature=dumpfile type=writer control=sequence
//ff:what newRunWriter — create a run file with a fast zstd encoder of a small window
package dumpfile

import (
	"os"

	"github.com/klauspost/compress/zstd"
)

// newRunWriter creates path (truncating a leftover of an interrupted run)
// with a zstd encoder at the fastest level, threads goroutines and a 1 MiB
// window, which keeps each decoder of a merge small (openRun).
func newRunWriter(path string, threads int) (*runWriter, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	z, err := zstd.NewWriter(f, zstd.WithEncoderLevel(zstd.SpeedFastest), zstd.WithEncoderConcurrency(max(1, threads)), zstd.WithWindowSize(1<<20))
	if err != nil {
		f.Close()
		return nil, err
	}
	return &runWriter{f: f, z: z}, nil
}
