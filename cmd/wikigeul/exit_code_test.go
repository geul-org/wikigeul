//ff:func feature=cli type=helper control=sequence
//ff:what TestExitCode — an anomaly stop is 3, a memory stop 4 with the resume command told, wrapped too; any other error is 1
package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/geul-org/wikigeul/pkg/memguard"
	"github.com/geul-org/wikigeul/pkg/pipeline"
)

func TestExitCode(t *testing.T) {
	var b bytes.Buffer
	if got := exitCode(fmt.Errorf("run: %w", pipeline.AnomalyError{}), nil, &b); got != 3 || b.Len() != 0 {
		t.Fatalf("anomaly: %d", got)
	}
	if got := exitCode(fmt.Errorf("run: %w", memguard.StopError{}), []string{"wikigeul", "-dump", "x y"}, &b); got != 4 || !strings.Contains(b.String(), "resume: wikigeul -dump 'x y' -resume") {
		t.Fatalf("memory: %d %q", got, b.String())
	}
	if got := exitCode(errors.New("x"), nil, &b); got != 1 {
		t.Fatalf("other: %d", got)
	}
}
