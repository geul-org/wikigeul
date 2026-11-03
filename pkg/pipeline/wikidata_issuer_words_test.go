//ff:func feature=pipeline type=helper control=sequence
//ff:what TestWikidataIssuerWords — wikidata.geul.org is the 2-word Issuer 0x00000001 (WikidataIssuer), canonical, a fresh slice each call
package pipeline

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestWikidataIssuerWords(t *testing.T) {
	a := wikidataIssuerWords()
	if !slices.Equal(a, []core.Word{0x0000, 0x0001}) || core.FormatID(a) != WikidataIssuer || WikidataIssuer != "0x00000001" || !core.IssuerCanonical(a) {
		t.Fatalf("%04X %s", a, WikidataIssuer)
	}
	a[1] = 9
	if wikidataIssuerWords()[1] != 1 {
		t.Fatal("shared slice")
	}
}
