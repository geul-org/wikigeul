//ff:func feature=pipeline type=command control=sequence
//ff:what TestPackWriterDrain — results in any order are put and the set closed; a failed result, a guard stop and a producer error end it
package pipeline

import (
	"errors"
	"testing"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

func TestPackWriterDrain(t *testing.T) {
	load := fixtureLoader(t)
	c := packTestConfig(t)
	s := packTestStream(t)
	feed := func(rs ...packResult) <-chan packResult {
		ch := make(chan packResult, len(rs))
		for _, r := range rs {
			ch <- r
		}
		close(ch)
		return ch
	}
	a := convertOne(load, c, packTestProv(c), s, PackJob{Seq: 1, Kind: 'Q', Num: 2})
	b := convertOne(load, c, packTestProv(c), s, PackJob{Seq: 0, Kind: 'Q', Num: 1})
	var none error
	w, _ := newPackWriter(s, Config{Every: 10}, PackProgress{})
	if err := w.drain(feed(a, b), &none); err != nil || w.prog.Docs != 2 || w.next != 2 {
		t.Fatal(err)
	}
	bad, _ := newPackWriter(packTestStream(t), Config{Every: 10}, PackProgress{})
	if err := bad.drain(feed(packResult{Err: errors.New("x")}), &none); err == nil {
		t.Fatal("failed result: want error")
	}
	stopped, _ := newPackWriter(packTestStream(t), Config{Every: 10, Guard: fixtureGuard(t, 10)}, PackProgress{})
	if err := stopped.drain(make(chan packResult), &none); !errors.As(err, &memguard.StopError{}) {
		t.Fatalf("guard stop: %v", err)
	}
	perr := errors.New("producer")
	p, _ := newPackWriter(packTestStream(t), Config{Every: 10}, PackProgress{})
	if err := p.drain(feed(), &perr); err != perr {
		t.Fatalf("producer error: %v", err)
	}
}
