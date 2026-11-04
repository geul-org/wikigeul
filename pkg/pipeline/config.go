//ff:type feature=pipeline type=model
//ff:what Config — what a conversion run shares between its inputs (the database through wikipgeul, or a dump file through wikigeul): archive root, QID range, commits, language registry, report, pack writing and strict mode
package pipeline

import (
	"github.com/geul-org/wikigeul/pkg/memguard"
	"github.com/geul-org/wikigeul/pkg/port"
)

// Config is one conversion run, whichever input it reads: the archive root
// (geul-dev draft) whose .well-known/geul receives
// the documents, the QID range [From, To] (every language is written: one
// document per entity holds them all), the porter and GEUL library commits
// recorded in TRANSFORMED-BY and where the report goes. Schema and Dump name
// the loaded dump in the checkpoint identity and in the report's first line
// (wikipg export-geul -schema, -dump-id); wikigeul keeps their defaults, so a
// checkpoint written before the two paths split still matches. LangRoot and
// LangVer say where and which version of the LANGUAGE-SHORT-CODE codec spec
// document to read the GEUL Language Registry from (LangRoot "" is Root,
// LangVer 0 the newest at the start of the run); LoadLanguages reads it once
// into Langs and LangDoc, so every document of a run uses the same registry.
// A run converts with Workers goroutines, the entities of IDs (a file of Q/P
// IDs) or else every property and Q<From>…Q<To>, checkpoints every Every
// documents, Resume continues an interrupted run, RecordedAt is the
// RECORDED-AT of every version list (0: the run's start) and Samples how many
// documents train the zstd dictionary. Strict stops the run at the first
// anomaly (anomaliesOf: a skip or raw-kept value outside the normal list, a
// failed conversion or panic, a missing entity, a reference mismatch, a
// schema Fact other than the known Warn E6), after a checkpoint, writing the
// anomaly report Anomaly; Allow adds "kind:reason" entries ("raw:time: 0년",
// "warn:E7") to the normal list for a resumed run. Source is not a flag: a
// dump read directly is named there by file name and SHA-256, part of the
// checkpoint identity, and empty for the loaded database. Mem are the memory
// limits of the -mem- flags and Guard the running watch over them (memguard.Start;
// nil: no guard); neither is part of the run's identity, since memory changes
// when work happens, not what is written.
type Config struct {
	Root       string
	From       uint64
	To         uint64
	Schema     string
	Dump       int16
	Commit     string
	LibCommit  string
	Report     string
	LangRoot   string
	LangVer    uint32
	Langs      port.Languages
	LangDoc    string
	Workers    int
	IDs        string
	Every      int64
	Resume     bool
	RecordedAt int64
	Samples    int
	Strict     bool
	Anomaly    string
	Allow      []string
	Source     string
	Mem        memguard.Limits
	Guard      *memguard.Guard
}
