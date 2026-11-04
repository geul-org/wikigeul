//ff:func feature=pipeline type=command control=sequence
//ff:what BindFlags — define the flags both conversion commands share and return the check to call after parsing
package pipeline

import (
	"flag"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

// BindFlags defines on fs the flags wikigeul and wikipg export-geul share,
// filling c, with -workers defaulting to workers and described by usage. The
// returned check (checkFlags) runs after fs.Parse.
func BindFlags(fs *flag.FlagSet, c *Config, workers int, usage string) func() error {
	fs.StringVar(&c.Root, "root", ".", "archive root holding .well-known/geul (geul-dev draft)")
	fs.Uint64Var(&c.From, "from", 1, "first QID number")
	fs.Uint64Var(&c.To, "to", 1000, "last QID number")
	fs.StringVar(&c.Commit, "commit", "", "wikigeul (converter) git commit, full ID (required)")
	fs.StringVar(&c.LibCommit, "lib-commit", "", "GEUL library git commit, full ID (required)")
	fs.StringVar(&c.Report, "report", "", "report TSV path (empty: none)")
	fs.StringVar(&c.LangRoot, "langreg-root", "", "archive root to read the LANGUAGE-SHORT-CODE spec (GEUL Language Registry) from (empty: -root)")
	fs.IntVar(&c.Workers, "workers", workers, usage)
	fs.StringVar(&c.IDs, "ids", "", "pack mode: file of entity IDs (Q…/P…, one per line) to convert instead of every property and the Q range")
	fs.Int64Var(&c.Every, "checkpoint", 100000, "pack mode: documents between checkpoints")
	fs.BoolVar(&c.Resume, "resume", false, "pack mode: continue the interrupted run in -root from its last checkpoint")
	fs.Int64Var(&c.RecordedAt, "recorded-at", 0, "pack mode: RECORDED-AT of the version lists, Unix seconds (0: when the run starts)")
	fs.IntVar(&c.Samples, "samples", 2048, "pack mode: documents the zstd dictionary is trained on (the first Q documents of the run)")
	fs.BoolVar(&c.Strict, "strict", false, "pack mode: stop at the first anomaly (a skip or raw value outside the normal list, an error, a schema Fact) after a checkpoint, exit 3")
	fs.StringVar(&c.Anomaly, "anomaly", "wikigeul-anomaly.tsv", "pack mode, -strict: where the anomaly report goes")
	fs.Func("allow", "pack mode, -strict: also treat kind:reason as normal (skipped:…, raw:…, warn:<rule>); repeatable", func(v string) error {
		c.Allow = append(c.Allow, v)
		return nil
	})
	var langVer uint
	fs.UintVar(&langVer, "langreg-version", 0, "version of the LANGUAGE-SHORT-CODE spec document to use (0: the newest when the run starts)")
	checkMem := memguard.BindFlags(fs, &c.Mem)
	return func() error {
		if err := checkMem(); err != nil {
			return err
		}
		return checkFlags(c, langVer)
	}
}
