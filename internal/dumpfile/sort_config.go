//ff:type feature=dumpfile type=model
//ff:what SortConfig — how to sort a dump into runs: the dump, the directory, memory per run, threads, which entities to keep, where to log
package dumpfile

import (
	"io"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

// SortConfig is one sort (Sort): the dump Info, the directory that holds
// the state and the runs, the bytes of entity JSON held before a run is
// written (Mem), threads for the decompressor and the zstd encoder, Filter
// (a text naming the kept entities, part of the state's identity) and Keep,
// which decides from the ID's letter and number whether an entity line is
// kept. Log receives progress and the lines that could not be read. Guard is
// the memory watch (nil: none): it shrinks the records held and the merges,
// and can stop the sort at a checkpoint.
type SortConfig struct {
	Info    Info
	Dir     string
	Mem     int64
	Threads int
	Filter  string
	Keep    func(kind byte, num uint64) bool
	Log     io.Writer
	Guard   *memguard.Guard
}
