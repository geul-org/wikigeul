//ff:type feature=cli type=model
//ff:what config — a wikigeul run: what every conversion run shares, and the dump file it reads with its provenance overrides and sort
package main

import "github.com/geul-org/wikigeul/pkg/pipeline"

// config is one wikigeul run, always reading a Wikidata JSON dump directly
// (pack mode): the shared run settings, then DumpFile (-dump), DumpURL,
// DumpSHA256 and DumpTime overriding the provenance taken from the file
// (dumpfile.InfoOf), SortDir holding the sort into key order and its
// checkpoints, SortMem the MiB of entity JSON per sorted run.
type config struct {
	pipeline.Config
	DumpFile   string
	DumpURL    string
	DumpSHA256 string
	DumpTime   int64
	SortDir    string
	SortMem    int64
}
