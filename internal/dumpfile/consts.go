//ff:type feature=dumpfile type=model
//ff:what consts — run file record header size, runs merged at once, dump text between sort checkpoints
package dumpfile

const (
	// recordHeaderSize is the fixed part of a record in a run file: kind
	// (1 byte), number (8), dump line (8) and data length (4), big-endian.
	recordHeaderSize = 21
	// mergeFanIn is how many runs are merged at once at most (reduceRuns, the
	// final merge): each open run holds a zstd decoder and a read buffer.
	mergeFanIn = 64
	// minFanIn is the narrowest merge the guard asks for: below it the runs
	// pass through too many merges.
	minFanIn = 8
	// openRunMemory is what one open run of a merge holds: its zstd decoder
	// with a 1 MiB window and its read buffer.
	openRunMemory = 3 << 20
	// sortBudgetPart is the part of the process's memory budget the records
	// held may take.
	sortBudgetPart = 4
	// checkpointBytes is how much dump text is read at most between sort
	// checkpoints when few lines are kept (a small -ids list).
	checkpointBytes = 16 << 30
)
