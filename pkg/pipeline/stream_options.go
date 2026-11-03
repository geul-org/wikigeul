//ff:func feature=pipeline type=helper control=sequence
//ff:what streamOptions — the pack stream's options for run c: zstd with one encoder per active worker, Close paced by the guard, its merge width fitted to the memory there is
package pipeline

import "github.com/geul-org/gogeul/pkg/pack"

// minMergeRuns is the narrowest merge width the guard asks of Close: a merge
// at two runs at a time would pass over a big run list too many times.
const minMergeRuns = 8

// streamOptions builds the options. Encoders stay allocated once used, and
// are used in turn, so fewer active workers need fewer encoders
// (Zstd.SetWorkers). With no guard the hooks change nothing: Pace goes on and
// MergeWidth is the default.
func streamOptions(c Config) pack.StreamOptions {
	z := &pack.Zstd{Workers: c.Workers}
	c.Guard.OnActive(z.SetWorkers)
	return pack.StreamOptions{Build: pack.BuildOptions{Codec: z}, Issuers: []string{WikidataIssuer}, Created: c.RecordedAt, Pace: c.Guard.Pace,
		MergeWidth: func() int { return c.Guard.Fit(pack.MergeRunMemory, minMergeRuns, pack.DefaultMergeRuns) }}
}
