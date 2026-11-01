//ff:func feature=pipeline type=command control=iteration dimension=1
//ff:what convertWorker — convert jobs until the channel closes, each result to out
package pipeline

import (
	"sync"

	"github.com/geul-org/gogeul/pkg/pack"

	"github.com/geul-org/wikigeul/pkg/port"
)

// convertWorker runs convertOne with its own loader for each job and
// signals wg when jobs is closed and drained. Each conversion waits at the
// guard's gate, so while memory is short fewer run at once, and ends the worker
// when the run is stopping. Only an entered conversion holds working memory; a
// waiting job holds its line, as it did in the channel.
func convertWorker(load Loader, c Config, prov port.Provenance, s *pack.Stream, jobs <-chan PackJob, out chan<- packResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		if !c.Guard.Enter() {
			return
		}
		r := convertOne(load, c, prov, s, j)
		c.Guard.Leave()
		out <- r
	}
}
