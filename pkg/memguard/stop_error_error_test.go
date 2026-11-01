//ff:func feature=memguard type=helper control=iteration dimension=1
//ff:what TestStopErrorError — the message names the phase, the reason and -resume
package memguard

import (
	"strings"
	"testing"
)

func TestStopErrorError(t *testing.T) {
	got := StopError{Phase: "close:h", Reason: "system free 9%"}.Error()
	for _, want := range []string{"close:h", "system free 9%", "-resume"} {
		if !strings.Contains(got, want) {
			t.Fatalf("%q lacks %q", got, want)
		}
	}
}
