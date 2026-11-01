//ff:func feature=pipeline type=command control=sequence
//ff:what dictSamples — the documents a pack-mode run trains its zstd dictionary on: the first Q entities of the run, built as they will be written
package pipeline

import "github.com/geul-org/wikigeul/pkg/port"

// dictSamples builds the documents of the run's sample entities, which each
// input lists itself (the first c.Samples Q entities of the ID list, FirstQ,
// else of the input from Q<From>) — items are almost all of Wikidata, and the
// same input always gives the same samples, so the same dictionary
// (pack.StreamOptions.Samples).
func dictSamples(load Loader, c Config, prov port.Provenance, samples func() ([]PackJob, error)) ([][]byte, error) {
	qs, err := samples()
	if err != nil {
		return nil, err
	}
	return buildAll(load, c, prov, qs)
}
