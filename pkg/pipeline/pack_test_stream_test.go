//ff:func feature=pipeline type=command control=sequence
//ff:what packTestStream — a new pack stream for wikidata.geul.org in a temp folder, Stored records (fast)
package pipeline

import (
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

// packTestStream starts a stream of Issuer 0x00000001 (wikidata.geul.org) without compression.
func packTestStream(t *testing.T) *pack.Stream {
	s, err := pack.NewStream(filepath.Join(t.TempDir(), "set"), pack.StreamOptions{Build: pack.BuildOptions{Codec: pack.Stored{}}, Issuers: []string{WikidataIssuer}, Created: 1})
	if err != nil {
		t.Fatal(err)
	}
	return s
}
