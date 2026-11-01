//ff:func feature=cli type=command control=sequence
//ff:what TestDecodeJob — a line decodes as its entity; a line that does not decode or names another entity is a dropError
package main

import (
	"errors"
	"testing"
)

func TestDecodeJob(t *testing.T) {
	if e, ok, err := decodeJob(packJob{Kind: 'P', Num: 31, Line: `{"id":"P31","datatype":"url"}`}); err != nil || !ok || e.Datatype != "url" {
		t.Fatalf("%+v %v %v", e, ok, err)
	}
	var de dropError
	if _, ok, err := decodeJob(packJob{Kind: 'Q', Num: 1, Line: `{"id":"Q1","claims":[]}`}); ok || !errors.As(err, &de) || de.ID != "Q1" {
		t.Fatalf("not decoded: %v", err)
	}
	if _, ok, err := decodeJob(packJob{Kind: 'Q', Num: 1, Line: `{"id":"Q2"}`}); ok || !errors.As(err, &de) {
		t.Fatalf("other entity: %v", err)
	}
}
