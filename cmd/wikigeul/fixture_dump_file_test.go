//ff:func feature=cli type=helper control=sequence
//ff:what fixtureDumpFile — a small Wikidata JSON dump file with its .sha256 sidecar in a temporary directory
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"
)

// testDumpBody is a dump in no ID order: Q42 (label, P31 → Q5 with a
// qualifier and a reference, sitelink), the property P31, Q1, a Lexeme and
// a line that does not decode as its entity (Q7, labels as an array).
const testDumpBody = "[\n" +
	`{"type":"item","id":"Q42","lastrevid":5,"modified":"2026-10-17T04:21:25Z","labels":{"en":{"language":"en","value":"Douglas Adams"}},"descriptions":{},"aliases":{},` +
	`"claims":{"P31":[{"mainsnak":{"snaktype":"value","property":"P31","datavalue":{"value":{"entity-type":"item","numeric-id":5,"id":"Q5"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q42$a","rank":"normal",` +
	`"qualifiers":{"P580":[{"snaktype":"novalue","property":"P580","datatype":"time"}]},"qualifiers-order":["P580"],` +
	`"references":[{"hash":"h","snaks":{"P248":[{"snaktype":"value","property":"P248","datavalue":{"value":{"entity-type":"item","numeric-id":9,"id":"Q9"},"type":"wikibase-entityid"},"datatype":"wikibase-item"}]},"snaks-order":["P248"]}]}]},` +
	`"sitelinks":{"enwiki":{"site":"enwiki","title":"Douglas Adams","badges":[]}}},` + "\n" +
	`{"type":"property","id":"P31","datatype":"wikibase-item","labels":{"en":{"language":"en","value":"instance of"}},"claims":{}},` + "\n" +
	`{"type":"item","id":"Q1","labels":{"en":{"language":"en","value":"Universe"}}},` + "\n" +
	`{"type":"lexeme","id":"L1"},` + "\n" +
	`{"type":"item","id":"Q7","labels":[]}` + "\n" +
	"]\n"

// fixtureDumpFile writes body as wikidata-20261019-all.json with its SHA-256
// in the .sha256 sidecar and returns the path.
func fixtureDumpFile(t *testing.T, body string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "wikidata-20261019-all.json")
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256([]byte(body))
	os.WriteFile(path+".sha256", []byte(hex.EncodeToString(sum[:])+"  wikidata-20261019-all.json\n"), 0o644)
	return path
}
