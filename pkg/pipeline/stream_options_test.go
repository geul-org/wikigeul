//ff:func feature=pipeline type=helper control=sequence
//ff:what TestStreamOptions — no guard: full merge width, Pace goes on; a short guard narrows the merge, stops Close and sizes the encoders by the active workers
package pipeline

import (
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

func TestStreamOptions(t *testing.T) {
	o := streamOptions(Config{Workers: 4, RecordedAt: 7})
	if o.MergeWidth() != pack.DefaultMergeRuns || o.Pace("close:h") != nil || o.Created != 7 || o.Issuers[0] != "0x00000001" {
		t.Fatalf("%+v", o)
	}
	if z, ok := o.Build.Codec.(*pack.Zstd); !ok || z.Workers != 4 {
		t.Fatalf("%+v", o.Build.Codec)
	}
	g := fixtureGuard(t, 20)
	s := streamOptions(Config{Workers: 8, Guard: g})
	if got := s.MergeWidth(); got != minMergeRuns {
		t.Fatalf("width %d", got)
	}
	z := s.Build.Codec.(*pack.Zstd)
	g.Pace("close:h")
	if z.Workers >= 8 {
		t.Fatalf("encoders %d", z.Workers)
	}
	var stop memguard.StopError
	if err := streamOptions(Config{Workers: 1, Guard: fixtureGuard(t, 10)}).Pace("close:h"); !errors.As(err, &stop) {
		t.Fatal(err)
	}
}
