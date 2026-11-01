//ff:func feature=pipeline type=command control=sequence
//ff:what TestConvertWorkerGate — a stopped guard ends the worker before it converts; under caution jobs still run, fewer at a time
package pipeline

import (
	"sync"
	"testing"
)

func TestConvertWorkerGate(t *testing.T) {
	load := fixtureLoader(t)
	c := packTestConfig(t)
	c.Guard = fixtureGuard(t, 10)
	jobs, out := make(chan PackJob, 1), make(chan packResult, 1)
	jobs <- PackJob{Kind: 'Q', Num: 1}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	convertWorker(load, c, packTestProv(c), packTestStream(t), jobs, out, wg)
	wg.Wait()
	if len(out) != 0 {
		t.Fatal("converted after the stop")
	}
	c.Guard = fixtureGuard(t, 20)
	jobs <- PackJob{Kind: 'Q', Num: 1}
	close(jobs)
	wg.Add(1)
	convertWorker(load, c, packTestProv(c), packTestStream(t), jobs, out, wg)
	if r := <-out; r.Err != nil {
		t.Fatal(r.Err)
	}
}
