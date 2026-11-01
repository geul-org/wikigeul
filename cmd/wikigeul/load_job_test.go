//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestLoadJob — a dump job decodes its line; one without a line is missing; an undecodable or mismatched line is a dropError
package main

import (
	"errors"
	"testing"
)

func TestLoadJob(t *testing.T) {
	e, ok, err := loadJob(nil, config{}, packJob{Kind: 'Q', Num: 42, Dump: true, Line: `{"type":"item","id":"Q42","labels":{"en":{"language":"en","value":"x"}}}`})
	if err != nil || !ok || e.ID != "Q42" || len(e.Labels) != 1 {
		t.Fatalf("%+v %v %v", e, ok, err)
	}
	if _, ok, err := loadJob(nil, config{}, packJob{Kind: 'Q', Num: 42, Dump: true}); ok || err != nil {
		t.Fatal("missing")
	}
	if _, _, err := loadJob(nil, config{Schema: "Bad-Schema"}, packJob{Kind: 'Q', Num: 42}); err == nil {
		t.Fatal("database job with a bad schema: want an error")
	}
	for _, line := range []string{`{"type":"item","id":"Q43"}`, `{"id":"Q42","labels":[]}`} {
		if _, _, err := loadJob(nil, config{}, packJob{Kind: 'Q', Num: 42, Dump: true, Line: line}); !errors.As(err, &dropError{}) {
			t.Errorf("%s: %v", line, err)
		}
	}
}
