//ff:func feature=pipeline type=command control=sequence
//ff:what TestPackWriterDrained — workers ended by the guard stop the run, a producer error fails it, otherwise the stream closes
package pipeline

import (
	"errors"
	"testing"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

func TestPackWriterDrained(t *testing.T) {
	var none error
	w, _ := newPackWriter(packTestStream(t), Config{Every: 10, Guard: fixtureGuard(t, 10)}, PackProgress{})
	var stop memguard.StopError
	if err := w.drained(&none); !errors.As(err, &stop) {
		t.Fatal(err)
	}
	perr := errors.New("producer")
	p, _ := newPackWriter(packTestStream(t), Config{Every: 10}, PackProgress{})
	if err := p.drained(&perr); err != perr {
		t.Fatal(err)
	}
	e, _ := newPackWriter(packTestStream(t), Config{Every: 10}, PackProgress{})
	if err := e.drained(&none); err == nil {
		t.Fatal("no document: want the Close error")
	}
}
