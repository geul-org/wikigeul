//ff:func feature=pipeline type=command control=iteration dimension=1
//ff:what buildAll — the documents of jobs (buildOne), skipping entities the dump lacks
package pipeline

import (
	"github.com/geul-org/wikigeul/pkg/port"
)

// buildAll builds each job's document in order; entities the dump does not
// have give none. The first error stops it.
func buildAll(load Loader, c Config, prov port.Provenance, js []PackJob) ([][]byte, error) {
	var out [][]byte
	for _, j := range js {
		b, _, ok, err := buildOne(load, c, prov, j)
		if err != nil {
			return nil, err
		}
		if ok {
			out = append(out, b)
		}
	}
	return out, nil
}
