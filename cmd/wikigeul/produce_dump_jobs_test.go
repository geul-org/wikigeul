//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestProduceDumpJobs — jobs after the checkpoint numbered from 0 in key order, the channel closed; missing runs leave the error
package main

import (
	"testing"

	"github.com/geul-org/wikigeul/internal/dumpfile"
)

func TestProduceDumpJobs(t *testing.T) {
	c, st := fixtureSortedDump(t, nil)
	jobs := make(chan packJob, 10)
	var err error
	produceDumpJobs(c, st, packProgress{Kind: 'Q', Num: 1}, nil, jobs, &err)
	var got []packJob
	for j := range jobs {
		got = append(got, j)
	}
	if err != nil || len(got) != 2 || got[0].Num != 7 || got[0].Seq != 0 || got[1].Num != 42 || got[1].Seq != 1 || got[1].Line == "" {
		t.Fatalf("%+v %v", got, err)
	}
	jobs = make(chan packJob, 10)
	produceDumpJobs(c, dumpfile.State{Runs: []string{"none"}}, packProgress{}, nil, jobs, &err)
	if _, open := <-jobs; open || err == nil {
		t.Fatal("missing run: want an error and a closed channel")
	}
}
