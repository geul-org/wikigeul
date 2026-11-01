//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestLoadJob — a job decodes its line; one without a line is missing; an undecodable or mismatched line is a DropError
package main

import (
	"errors"
	"testing"

	"github.com/geul-org/wikigeul/pkg/pipeline"
)

func TestLoadJob(t *testing.T) {
	e, ok, err := loadJob(pipeline.PackJob{Kind: 'Q', Num: 42, Line: `{"type":"item","id":"Q42","labels":{"en":{"language":"en","value":"x"}}}`})
	if err != nil || !ok || e.ID != "Q42" || len(e.Labels) != 1 {
		t.Fatalf("%+v %v %v", e, ok, err)
	}
	if _, ok, err := loadJob(pipeline.PackJob{Kind: 'Q', Num: 42}); ok || err != nil {
		t.Fatal("missing")
	}
	for _, line := range []string{`{"type":"item","id":"Q43"}`, `{"id":"Q42","labels":[]}`} {
		if _, _, err := loadJob(pipeline.PackJob{Kind: 'Q', Num: 42, Line: line}); !errors.As(err, &pipeline.DropError{}) {
			t.Errorf("%s: %v", line, err)
		}
	}
}
