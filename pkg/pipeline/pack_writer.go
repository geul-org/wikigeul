//ff:type feature=pipeline type=model
//ff:what packWriter — the pack-mode writer: stream, report file, progress, the expected TRANSFORMED-BY and the results waiting for their turn
package pipeline

import (
	"bufio"
	"os"
	"time"

	"github.com/geul-org/gogeul/pkg/pack"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

// packWriter puts results in job order (take) into the stream s, appends
// their report lines to rep (buffered in w), checkpoints every `every`
// documents with prog, and checks that every document carries the
// TRANSFORMED-BY records tb of the first one. A strict writer stops at the
// first result with anomalies (halt), writing them to anomaly with the run
// line run and the anomalies allowed by -allow. skipped and raw sum the
// reasons of the documents put since the process started (from docs0 at
// start), printed with each checkpoint (logProgress). guard is the memory
// watch: the writer checkpoints when it enters caution (cautions counts how
// many entries it has answered) and stops when it says so (memStop).
type packWriter struct {
	s       *pack.Stream
	rep     *os.File
	w       *bufio.Writer
	prog    PackProgress
	every   int64
	since   int64
	tb      string
	next    int64
	pending map[int64]packResult
	strict  bool
	allow   map[string]bool
	anomaly string
	run     string
	skipped map[string]int
	raw     map[string]int
	start   time.Time
	docs0   int64
	guard   *memguard.Guard

	cautions int64
}
