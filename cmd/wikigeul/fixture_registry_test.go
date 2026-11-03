//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what fixtureRegistry — publish a small LANGUAGE-SHORT-CODE codec spec document (en, pt-BR, zh-Hans) into root's data folder
package main

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// fixtureRegistry publishes, as version 1 of the codec 0x00000000/codecs/0x000007a0, a
// codec document like the real one: DOC-ID CODEC-POINTER(0x00000000,
// 0x000007a0, 0), DECLARE 0x07A0 (LANGUAGE-SHORT-CODE), the 1-WORD FIELD code and
// one ENUM with its BIND per tag, here en 0x0784, pt-BR 0x4130, zh-Hans 0x404c
// (only those in tags when any are given; tags naming none gives a spec
// without codes).
func fixtureRegistry(t *testing.T, root string, tags ...string) {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	bind := func(id, subj uint64, name string) []core.Word {
		r, _ := core.EncodeBind(core.Bind{EdgeID: id, Subject: core.EncodeLocalRef(subj), Name: []byte(name)})
		return enc(r)
	}
	target := core.TargetWord(core.KindValue, core.CodecLanguageShortCode)
	docID, _ := core.EncodeDocIDCodec([]core.Word{0, 0}, uint32(target))
	in := slices.Concat(enc(core.EncodeSpecVersion(1)), enc(docID), []core.Word{core.TargetWord(core.KindMeta, core.MetaBody)},
		enc(core.EncodeDeclare(core.Declare{EdgeID: 1, Target: target})),
		enc(core.EncodeFieldDecl(core.FieldDecl{EdgeID: 2, Declare: 1, Width: 1})),
		bind(3, 1, "LANGUAGE-SHORT-CODE"), bind(4, 2, "code"))
	id := uint64(5)
	for _, e := range []struct {
		code core.Word
		tag  string
	}{{0x0784, "en"}, {0x404C, "zh-Hans"}, {0x4130, "pt-BR"}} {
		if len(tags) > 0 && !slices.Contains(tags, e.tag) {
			continue
		}
		in = append(in, enc(core.EncodeEnumDecl(core.EnumDecl{EdgeID: id, Field: 2, Code: []core.Word{e.code}}))...)
		in = append(in, bind(id+1, id, e.tag)...)
		id += 2
	}
	b := core.PutWords(core.WrapDocument(append(in, core.TargetWord(core.KindMeta, core.MetaEnd))))
	if _, err := (resolve.Dir{Root: root}).PublishCodec(b); err != nil {
		t.Fatal(err)
	}
}
