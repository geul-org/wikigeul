//ff:func feature=cli type=helper control=sequence
//ff:what dumpProvenance — the provenance every document of a run reading a dump directly records: the dump's URL, SHA-256 and time, the converter and library commits
package main

import (
	"github.com/geul-org/wikigeul/internal/dumpfile"
	"github.com/geul-org/wikigeul/pkg/pipeline"
	"github.com/geul-org/wikigeul/pkg/port"
)

// dumpProvenance is the provenance of a dump read directly: the URL, file
// SHA-256 and download time come from d (dumpfile.InfoOf, the values wikipgeul
// would have put in the dump row), the tools are the same two (wikigeul at
// commit, gogeul at libCommit: pipeline.Tools). No loader is involved.
func dumpProvenance(d dumpfile.Info, commit, libCommit string) port.Provenance {
	return port.Provenance{
		DumpURI:    d.URL,
		DumpSHA256: d.SHA256,
		AcquiredAt: d.DownloadedAt,
		Tools:      pipeline.Tools(commit, libCommit),
	}
}
