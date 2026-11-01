//ff:func feature=pipeline type=command control=iteration dimension=1
//ff:what TestConvertJobs — two workers convert every job once and the channel closes
package pipeline

import (
	"testing"
)

func TestConvertJobs(t *testing.T) {
	load := fixtureLoader(t)
	c := packTestConfig(t)
	jobs := make(chan PackJob, 4)
	for i, n := range []uint64{1, 2, 42} {
		jobs <- PackJob{Seq: int64(i), Kind: 'Q', Num: n}
	}
	close(jobs)
	seen := map[int64]bool{}
	for r := range convertJobs([]Loader{load, load}, c, packTestProv(c), packTestStream(t), jobs) {
		if r.Err != nil || seen[r.Job.Seq] {
			t.Fatalf("%+v", r.Err)
		}
		seen[r.Job.Seq] = true
	}
	if len(seen) != 3 {
		t.Fatal(seen)
	}
}
