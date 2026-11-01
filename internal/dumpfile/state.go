//ff:type feature=dumpfile type=model
//ff:what State — the sort of one dump into runs: which dump and filter, the runs, how far the dump was read, whether it is done
package dumpfile

// State is saved (state.json in the sort directory) at every sort
// checkpoint. Name, Size, ModTime, SHA256 and Filter say which dump and which
// entities (the run's range or ID list) it sorts; a resumed sort must match
// them (sameSort). Runs are the run files in the directory, Next numbers the
// next one. Lines and Offset are the dump lines and decompressed bytes read
// so far, every kept line of which is in Runs; Hash is the SHA-256 state of
// a plain file at Offset (a compressed file is hashed again from its start
// when resumed). Kept, Skipped and Bad count entity lines kept, left out by
// the filter or another kind (L, E …), and not read (no id, an ID wikipgeul
// cannot split). Scanned is set when the whole file was read and its SHA-256
// matched; Done when the runs were merged down to at most mergeFanIn.
type State struct {
	Name    string
	Size    int64
	ModTime int64
	SHA256  string
	Filter  string
	Runs    []string
	Next    int
	Lines   int64
	Offset  int64
	Hash    []byte
	Kept    int64
	Skipped int64
	Bad     int64
	Scanned bool
	Done    bool
}
