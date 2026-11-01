//ff:type feature=pipeline type=model
//ff:what PackProgress — what a pack-mode run saves with each stream checkpoint: the run's identity and how far it got
package pipeline

// PackProgress is saved (JSON) with every stream checkpoint
// (pack.Stream.Checkpoint meta). The identity fields — porter and library
// commits (TRANSFORMED-BY), registry version, dump, range or ID list, the
// RECORDED-AT of the version lists, the sampled dictionary — must match on
// resume, or the resumed documents would differ from the ones already
// written. Source names a dump read directly (-dump: file name and SHA-256)
// and is empty for the loaded database, so checkpoints written before it
// existed still match. Kind/Num is the last entity put, Docs how many documents, Report
// the report file's size at the checkpoint. Drained is set by the checkpoint
// made just before the stream closes: every entity is in, so a resumed run
// has none left to list.
type PackProgress struct {
	Commit     string
	LibCommit  string
	LangDoc    string
	Schema     string
	Dump       int16
	From       uint64
	To         uint64
	IDs        string
	Source     string `json:",omitempty"`
	RecordedAt int64
	Kind       byte
	Num        uint64
	Docs       int64
	Report     int64
	Drained    bool `json:",omitempty"`
}
