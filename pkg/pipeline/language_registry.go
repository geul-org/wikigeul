//ff:func feature=pipeline type=loader control=sequence
//ff:what languageRegistry — the GEUL Language Registry from one version of the LANGUAGE-SHORT-CODE codec spec document, with the version and content hash it came from
package pipeline

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/resolve"

	"github.com/geul-org/wikigeul/pkg/port"
)

// registryDoc is the codec ID of the LANGUAGE-SHORT-CODE codec document of the
// GEUL Issuer 0x00000000: its TAG with L=0 as 2 words, 0x000007A0
// (§3.13 `0x00000000/codecs/<코덱 ID>/`).
var registryDoc = []core.Word{0, core.TargetWord(core.KindValue, core.CodecLanguageShortCode)}

// languageRegistry reads version v of the LANGUAGE-SHORT-CODE codec document
// from dir (v 0: the newest neither deprecated nor deleted, fixed once here
// for the whole run) through the GEUL library's resolve path, and returns the
// registry its ENUMs and BINDs declare (codecspec.Spec.Languages, registryOf)
// and the version it read as `0x00000000/codecs/0x000007A0@n 0x<content
// hash>`, which the run prints and the report records. A document without
// registry is an error.
func languageRegistry(dir resolve.Dir, v uint32) (port.Languages, string, error) {
	where := core.FormatID(geulIssuerWords()) + "/codecs/" + core.FormatID(registryDoc)
	h, err := dir.Log(geulIssuerWords(), registryDoc)
	if err != nil {
		return nil, "", fmt.Errorf("language registry: %w", err)
	}
	e, ok := registryVersion(h, v)
	if !ok {
		return nil, "", fmt.Errorf("language registry: no version %d of %s", v, where)
	}
	b, err := dir.At(geulIssuerWords(), registryDoc, e.Version)
	if err != nil {
		return nil, "", fmt.Errorf("language registry: %w", err)
	}
	doc, _ := inspect.Decode(b) // At checked the content hash; a document that does not decode declares no code
	langs, err := registryOf(codecspec.Collect(bodyRecords(doc)).Languages())
	if err != nil {
		return nil, "", err
	}
	return langs, fmt.Sprintf("%s@%d 0x%X", where, e.Version, e.Hash), nil
}
