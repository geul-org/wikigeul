//ff:func feature=pipeline type=command control=iteration dimension=1
//ff:what convertWorker — convert jobs until the channel closes, each result to out
package pipeline

import (
	"sync"

	"github.com/geul-org/gogeul/pkg/pack"

	"github.com/geul-org/wikigeul/pkg/port"
)

// convertWorker runs convertOne with its own loader for each job and
// signals wg when jobs is closed and drained.
func convertWorker(load Loader, c Config, prov port.Provenance, s *pack.Stream, jobs <-chan PackJob, out chan<- packResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		out <- convertOne(load, c, prov, s, j)
	}
}
