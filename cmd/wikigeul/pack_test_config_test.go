//ff:func feature=cli type=command control=sequence
//ff:what packTestConfig — a pack-mode config: small checkpoints, two workers, fixed RECORDED-AT, a temp root
package main

import (
	"strings"
	"testing"

	"github.com/geul-org/wikigeul/pkg/pipeline"
)

// packTestConfig is a pack-mode run into a temp root.
func packTestConfig(t *testing.T) config {
	a := strings.Repeat("a", 40)
	return config{Config: pipeline.Config{Root: t.TempDir(), From: 1, To: 100, Schema: "wikidata", Dump: 1, Commit: a, LibCommit: a,
		Workers: 2, Every: 2, Samples: 2, RecordedAt: 1_800_000_000}}
}
