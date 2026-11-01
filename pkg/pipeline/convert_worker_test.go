//ff:func feature=pipeline type=command control=sequence
//ff:what TestConvertWorker — each job's result goes out and the group is told when the jobs end
package pipeline

import (
	"sync"
	"testing"
)

func TestConvertWorker(t *testing.T) {
	load := fixtureLoader(t)
	c := packTestConfig(t)
	jobs, out := make(chan PackJob, 1), make(chan packResult, 1)
	jobs <- PackJob{Kind: 'Q', Num: 1}
	close(jobs)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	convertWorker(load, c, packTestProv(c), packTestStream(t), jobs, out, wg)
	wg.Wait()
	if r := <-out; r.Err != nil || r.Doc != 0x5100000000000001 {
		t.Fatalf("%+v", r.Err)
	}
}
