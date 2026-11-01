//ff:func feature=cli type=helper control=sequence
//ff:what TestDumpProvenance — the dump file's URL, SHA-256 and time with the two tools at their commits, as provenanceOf gives for the loaded dump
package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/geul-org/wikigeul/internal/dumpfile"
	"github.com/geul-org/wikigeul/internal/source"
)

func TestDumpProvenance(t *testing.T) {
	got := dumpProvenance(dumpfile.Info{URL: "u", SHA256: "ab", DownloadedAt: 7}, "c", "l")
	want, _ := provenanceOf(source.Dump{URL: "u", SHA256: "ab", DownloadedAt: time.Unix(7, 0), Loader: wikipgPath}, "c", "l")
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%+v\n%+v", got, want)
	}
}
