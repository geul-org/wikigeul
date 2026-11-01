//ff:func feature=pipeline type=command control=sequence
//ff:what TestWritePackDrained — a run resumed with every entity in does not call the producer and goes straight to Close
package pipeline

import (
	"errors"
	"testing"
)

func TestWritePackDrained(t *testing.T) {
	load := fixtureLoader(t)
	c := packTestConfig(t)
	produce := func(jobs chan<- PackJob, errp *error) {
		*errp = errors.New("listed again")
		close(jobs)
	}
	err := WritePack(packTestStream(t), c, PackProgress{Drained: true, Docs: 1, Kind: 'Q', Num: 1}, packTestProv(c), []Loader{load}, produce)
	if err == nil || err.Error() == "listed again" {
		t.Fatalf("want Close's no-record error, got %v", err)
	}
}
