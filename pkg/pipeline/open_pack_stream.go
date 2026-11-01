//ff:func feature=pipeline type=command control=sequence
//ff:what OpenPackStream — the Issuer's pack stream: resumed from its checkpoint (same run only) or new with a dictionary trained on samples
package pipeline

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/geul-org/gogeul/pkg/pack"

	"github.com/geul-org/wikigeul/pkg/port"
)

// OpenPackStream opens the stream for set. With -resume it continues the
// checkpointed stream; the checkpoint must come from the same run (sameRun:
// the same commits in TRANSFORMED-BY, registry, dump, range or list,
// RECORDED-AT) — a resumed run that differed would mix documents. Otherwise
// it starts a new set: zstd with Workers concurrent encoders and a
// dictionary trained on dictSamples (samples lists its entities, load reads
// them; neither is touched on resume), created at RECORDED-AT. The guard
// sizes the compressor's encoders by the active workers, paces Close and
// picks its merge width (streamOptions).
func OpenPackStream(set string, c Config, prov port.Provenance, load Loader, samples func() ([]PackJob, error), idsHash string) (*pack.Stream, PackProgress, error) {
	want := progressOf(c, idsHash)
	o := streamOptions(c)
	if c.Resume {
		s, err := pack.ResumeStream(set, o)
		if err != nil {
			return nil, want, err
		}
		var p PackProgress
		err = json.Unmarshal(s.Meta(), &p)
		if c.RecordedAt == 0 {
			want.RecordedAt = p.RecordedAt // -recorded-at left out: the checkpoint's
		}
		if err != nil || !sameRun(p, want) {
			return nil, want, errors.New("wikigeul: -resume: the checkpoint is from another run (commits, registry, dump, range, ID list or -recorded-at differ)")
		}
		return s, p, nil
	}
	c.Guard.SetPhase("dictionary")
	docs, err := dictSamples(load, c, prov, samples)
	if err != nil {
		return nil, want, fmt.Errorf("dictionary samples: %w", err)
	}
	o.Samples = docs
	s, err := pack.NewStream(set, o)
	return s, want, err
}
