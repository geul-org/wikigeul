//ff:func feature=cli type=command control=sequence
//ff:what sortDump — sort the run's entities of a dump into key-ordered runs in -sort-dir (every property and the Q range, or the ID list)
package main

import (
	"fmt"
	"os"

	"github.com/geul-org/wikigeul/internal/dumpfile"
)

// sortDump sorts the dump into runs under c.SortDir (dumpfile.Sort, resumed
// from its own checkpoint when run again), keeping the entities the run
// converts (keepEntity) — every property and Q<From>…Q<To>, or the ID list —
// with -sort-mem MiB of entity JSON per run and c.Workers threads. The
// filter text names them in the sort state.
func sortDump(c config, info dumpfile.Info, ids []packJob, idsHash string) (dumpfile.State, error) {
	filter := fmt.Sprintf("P all, Q%d-Q%d", c.From, c.To)
	if ids != nil {
		filter = "ids sha256:" + idsHash
	}
	return dumpfile.Sort(dumpfile.SortConfig{Info: info, Dir: c.SortDir, Mem: c.SortMem << 20, Threads: c.Workers,
		Filter: filter, Keep: keepEntity(c, ids), Log: os.Stdout})
}
