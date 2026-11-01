//ff:func feature=cli type=helper control=sequence
//ff:what TestDumpProvenance — the dump file's URL, SHA-256 and time with the two tools at their commits
package main

import (
	"reflect"
	"testing"

	"github.com/geul-org/wikigeul/internal/dumpfile"
	"github.com/geul-org/wikigeul/pkg/pipeline"
	"github.com/geul-org/wikigeul/pkg/port"
)

func TestDumpProvenance(t *testing.T) {
	got := dumpProvenance(dumpfile.Info{URL: "u", SHA256: "ab", DownloadedAt: 7}, "c", "l")
	want := port.Provenance{DumpURI: "u", DumpSHA256: "ab", AcquiredAt: 7, Tools: pipeline.Tools("c", "l")}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%+v\n%+v", got, want)
	}
}
