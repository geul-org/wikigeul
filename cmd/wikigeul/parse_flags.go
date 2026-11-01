//ff:func feature=cli type=command control=sequence
//ff:what parseFlags — the run config from command-line arguments; the dump file and the porter and library commits are required
package main

import (
	"flag"

	"github.com/geul-org/wikigeul/pkg/pipeline"
)

// parseFlags reads args into a config: the flags every conversion run shares
// (pipeline.BindFlags), then the dump file's (checkDumpFlags). Schema and
// Dump keep the values the checkpoint identity has always held for a run
// reading a dump (pipeline.Config).
func parseFlags(args []string) (config, error) {
	fs := flag.NewFlagSet("wikigeul", flag.ContinueOnError)
	c := config{Config: pipeline.Config{Schema: "wikidata", Dump: 1}}
	check := pipeline.BindFlags(fs, &c.Config, defaultDumpWorkers(), "conversion goroutines, also the sort's decompressor and encoder threads")
	fs.StringVar(&c.DumpFile, "dump", "", "Wikidata JSON dump to read: .json, .json.gz or .json.bz2 (required)")
	fs.StringVar(&c.DumpURL, "dump-url", "", "URL recorded as the dump's source (empty: dumps.wikimedia.org URL from the file name's date)")
	fs.StringVar(&c.DumpSHA256, "dump-sha256", "", "the file's SHA-256 in hex (empty: from <file>.sha256); the file is checked against it")
	fs.Int64Var(&c.DumpTime, "dump-time", 0, "when the file was downloaded, Unix seconds (0: its modification time)")
	fs.StringVar(&c.SortDir, "sort-dir", "wikigeul-dumpsort", "directory of the sort into document key order and its checkpoints")
	fs.Int64Var(&c.SortMem, "sort-mem", 1024, "MiB of entity JSON held per sorted run")
	if err := fs.Parse(args); err != nil {
		return config{}, err
	}
	if err := check(); err != nil {
		return config{}, err
	}
	if err := checkDumpFlags(c); err != nil {
		return config{}, err
	}
	return c, nil
}
