//ff:func feature=port type=model control=iteration dimension=1
//ff:what TestCodes — each code this porter writes is the §3 TAG ID (L=0) and names exactly one code of the 512-code TAG space
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCodes(t *testing.T) {
	// TAG ID (L=0) from migration1 §3.1, §3.4, §3.5: written here, not taken from core.
	want := []struct {
		name  string
		kind  core.Kind
		codec uint16
		tag   core.Word
	}{
		{"URI-EVEN", core.KindValue, codecURIEven, 0x0540},
		{"LANGUAGE-SHORT-CODE", core.KindValue, codecLanguageShort, 0x05A0},
		{"REFERENCE-NODE", core.KindEdge, edgeReferenceNode, 0x0800},
		{"LABEL", core.KindEdge, edgeLabel, 0x0840},
		{"ALIAS", core.KindEdge, edgeAlias, 0x0850},
		{"DESCRIPTION", core.KindEdge, edgeDescription, 0x0860},
		{"TRIPLE", core.KindEdge, edgeTriple, 0x08A0},
		{"QUALIFIER", core.KindEdge, edgeQualifier, 0x08B0},
		{"SOURCE", core.KindEdge, edgeSource, 0x0A00},
		{"TRANSFORMED-BY", core.KindMeta, metaTransformedBy, 0x1D40},
		{"EXTERNAL-ISSUER-2WORD", core.KindReference, toolRef(1).Codec, 0x02D0},
	}
	for _, w := range want {
		if got := core.TargetWord(w.kind, w.codec); got != w.tag {
			t.Errorf("%s: TAG 0x%04X want 0x%04X", w.name, got, w.tag)
		}
		if codes := tagCodesOf(w.kind, w.codec); len(codes) != 1 || core.Word(codes[0])<<4 != w.tag {
			t.Errorf("%s: codes %03X of the TAG space, want the one TAG 0x%04X", w.name, codes, w.tag)
		}
	}
}
