//ff:func feature=pipeline type=helper control=sequence
//ff:what fixtureLoader — test fixture: a Loader over a few entity lines (Q1, Q2, Q7, Q42); any other entity is missing
package pipeline

import (
	"strconv"
	"testing"

	"github.com/geul-org/wikigeul/pkg/source"
)

// fixtureLines are the entities of fixtureLoader: Q42 with a label, a
// statement, a qualifier, a reference and a sitelink.
var fixtureLines = map[string]string{
	"Q1": `{"type":"item","id":"Q1","lastrevid":1,"labels":{"en":{"language":"en","value":"Universe"}}}`,
	"Q2": `{"type":"item","id":"Q2","lastrevid":2,"labels":{"en":{"language":"en","value":"Earth"}}}`,
	"Q7": `{"type":"item","id":"Q7","lastrevid":7,"labels":{"en":{"language":"en","value":"seven"}}}`,
	"Q42": `{"type":"item","id":"Q42","lastrevid":5,"modified":"2026-10-17T04:21:25Z","labels":{"en":{"language":"en","value":"Douglas Adams"}},` +
		`"claims":{"P31":[{"mainsnak":{"snaktype":"value","property":"P31","datavalue":{"value":{"entity-type":"item","numeric-id":5,"id":"Q5"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q42$a","rank":"normal",` +
		`"qualifiers":{"P580":[{"snaktype":"novalue","property":"P580","datatype":"time"}]},"qualifiers-order":["P580"],` +
		`"references":[{"hash":"h","snaks":{"P248":[{"snaktype":"value","property":"P248","datavalue":{"value":{"entity-type":"item","numeric-id":9,"id":"Q9"},"type":"wikibase-entityid"},"datatype":"wikibase-item"}]},"snaks-order":["P248"]}]}]},` +
		`"sitelinks":{"enwiki":{"site":"enwiki","title":"Douglas Adams","badges":[]}}}`,
}

// fixtureLoader decodes job's line of fixtureLines (source.DecodeDumpEntity).
func fixtureLoader(t *testing.T) Loader {
	return func(job PackJob) (source.Entity, bool, error) {
		line, ok := fixtureLines[string(job.Kind)+strconv.FormatUint(job.Num, 10)]
		if !ok {
			return source.Entity{}, false, nil
		}
		e, err := source.DecodeDumpEntity([]byte(line))
		if err != nil {
			t.Fatal(err)
		}
		return e, true, nil
	}
}
