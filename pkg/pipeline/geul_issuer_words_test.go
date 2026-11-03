//ff:func feature=pipeline type=helper control=sequence
//ff:what TestGeulIssuerWords — geul.org is the 2-word Issuer 0x00000000, a fresh slice each call
package pipeline

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestGeulIssuerWords(t *testing.T) {
	a := geulIssuerWords()
	if !slices.Equal(a, []core.Word{0x0000, 0x0000}) || core.FormatID(a) != "0x00000000" || !core.IssuerCanonical(a) {
		t.Fatalf("%04X", a)
	}
	a[1] = 9
	if geulIssuerWords()[1] != 0 {
		t.Fatal("shared slice")
	}
}
