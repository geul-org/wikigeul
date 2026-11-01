//ff:func feature=pipeline type=command control=sequence
//ff:what TestConvertOne — a document with its list, both encoded, and its document number; a missing entity, a failing input and a dropped line carry their state; strict checks find nothing on a good document; a panic becomes the error
package pipeline

import (
	"errors"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
	"github.com/geul-org/wikigeul/pkg/source"
)

func TestConvertOne(t *testing.T) {
	load := fixtureLoader(t)
	c := packTestConfig(t)
	s := packTestStream(t)
	r := convertOne(load, c, packTestProv(c), s, PackJob{Kind: 'Q', Num: 42, Seq: 3})
	list, _ := resolve.FirstVersionList(r.Raw, c.RecordedAt)
	if r.Err != nil || r.Missing || r.Doc != 0x510000000000002a || string(r.List.Bytes) != string(list) || r.Version.Raw != len(r.Raw) || !r.Version.HashOK || r.Job.Seq != 3 {
		t.Fatalf("%+v", r.Err)
	}
	if r := convertOne(load, c, packTestProv(c), s, PackJob{Kind: 'Q', Num: 999}); !r.Missing || r.Err != nil {
		t.Fatal("missing")
	}
	strict := c
	strict.Strict = true
	if r := convertOne(load, strict, packTestProv(c), s, PackJob{Kind: 'Q', Num: 42}); r.Err != nil || r.Facts != nil && len(r.Facts) != 0 {
		t.Fatalf("strict: %v %+v", r.Err, r.Facts)
	}
	if r := convertOne(nil, c, packTestProv(c), s, PackJob{Kind: 'Q', Num: 42}); r.Err == nil || !strings.HasPrefix(r.Err.Error(), "panic: ") {
		t.Fatalf("nil loader: want a recovered panic, got %v", r.Err)
	}
	if r := convertOne(failLoader, c, packTestProv(c), s, PackJob{Kind: 'Q', Num: 42}); r.Err == nil {
		t.Fatal("failing input: want error")
	}
	dropped := func(PackJob) (source.Entity, bool, error) {
		return source.Entity{}, false, DropError{ID: "Q42", Err: errors.New("x")}
	}
	if r := convertOne(dropped, c, packTestProv(c), s, PackJob{Kind: 'Q', Num: 42}); r.Err != nil || r.Dropped == "" {
		t.Fatalf("dropped line: %+v", r)
	}
}
