//ff:func feature=pipeline type=helper control=sequence
//ff:what TestDropErrorError — the entity and the decoding error
package pipeline

import (
	"errors"
	"testing"
)

func TestDropErrorError(t *testing.T) {
	if got := (DropError{ID: "Q7", Err: errors.New("bad")}).Error(); got != "dump line of Q7 does not decode: bad" {
		t.Fatalf("%q", got)
	}
}
