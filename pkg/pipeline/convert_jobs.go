//ff:func feature=pipeline type=command control=iteration dimension=1
//ff:what convertJobs — start one worker per loader over the jobs; results arrive in any order and the channel closes when all are done
package pipeline

import (
	"sync"

	"github.com/geul-org/gogeul/pkg/pack"

	"github.com/geul-org/wikigeul/pkg/port"
)

// convertJobs runs convertWorker on each of loaders. The writer restores the
// output order from the job numbers (packWriter.take).
func convertJobs(loaders []Loader, c Config, prov port.Provenance, s *pack.Stream, jobs <-chan PackJob) <-chan packResult {
	out := make(chan packResult, 4*len(loaders))
	wg := &sync.WaitGroup{}
	for _, load := range loaders {
		wg.Add(1)
		go convertWorker(load, c, prov, s, jobs, out, wg)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
