//ff:func feature=pipeline type=command control=sequence
//ff:what TestPackWriterReceive — a result arrives; closed results say so; a guard stop ends the wait even with no result coming
package pipeline

import (
	"errors"
	"testing"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

func TestPackWriterReceive(t *testing.T) {
	ch := make(chan packResult, 1)
	ch <- packResult{Doc: 7}
	w := &packWriter{}
	if r, ok, err := w.receive(ch); err != nil || !ok || r.Doc != 7 {
		t.Fatalf("%+v %v %v", r, ok, err)
	}
	close(ch)
	if _, ok, err := w.receive(ch); err != nil || ok {
		t.Fatalf("%v %v", ok, err)
	}
	s, _ := newPackWriter(packTestStream(t), Config{Every: 10, Guard: fixtureGuard(t, 10)}, PackProgress{})
	var stop memguard.StopError
	if _, _, err := s.receive(make(chan packResult)); !errors.As(err, &stop) {
		t.Fatal(err)
	}
}
