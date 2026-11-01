//ff:func feature=pipeline type=helper control=sequence
//ff:what TestPackWriterCheckpointDue — due every documents; once for an entry into caution; not otherwise
package pipeline

import "testing"

func TestPackWriterCheckpointDue(t *testing.T) {
	w := &packWriter{every: 2}
	if w.checkpointDue() {
		t.Fatal("not yet")
	}
	w.since = 2
	if !w.checkpointDue() {
		t.Fatal("every")
	}
	w = &packWriter{every: 100, guard: fixtureGuard(t, 20)}
	if !w.checkpointDue() {
		t.Fatal("caution entered")
	}
	if w.checkpointDue() {
		t.Fatal("answered once")
	}
}
