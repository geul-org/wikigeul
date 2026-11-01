//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestRunDump — a dump file converts into the Issuer's pack set with its report, the undecodable line left out; strict mode stops at it; bad dump, registry, list or sort directory fail
package main

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestRunDump(t *testing.T) {
	dump := fixtureDumpFile(t, testDumpBody)
	reg := t.TempDir()
	fixtureRegistry(t, reg)
	a := strings.Repeat("a", 40)
	args := func(root string, extra ...string) []string {
		return append([]string{"-root", root, "-langreg-root", reg, "-from", "1", "-to", "100", "-commit", a, "-lib-commit", a, "-workers", "2",
			"-checkpoint", "1", "-recorded-at", "1800000000", "-samples", "2", "-dump", dump, "-sort-dir", filepath.Join(root, "sort")}, extra...)
	}
	root := t.TempDir()
	rep := filepath.Join(root, "r.tsv")
	if err := run(args(root, "-report", rep)); err != nil {
		t.Fatal(err)
	}
	r, err := pack.Open(context.Background(), pack.FileStore{Root: filepath.Join(root, ".well-known/geul/0x00010000")}, pack.Options{})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := r.Fetch(context.Background(), "0x00010000/0x510000000000002a/0x00000001.geul"); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(rep)
	if strings.Count(string(b), "\tok\t") != 3 || !strings.HasPrefix(string(b), "# wikigeul -from 1 -to 100 -dump wikidata-20261019-all.json sha256:") || strings.Contains(string(b), "Q7\t") {
		t.Fatalf("report:\n%s", b)
	}
	strict := t.TempDir()
	if err := run(args(strict, "-strict", "-anomaly", filepath.Join(strict, "a.tsv"))); !errors.As(err, &anomalyError{}) {
		t.Fatalf("strict: %v", err)
	}
	if err := run(args(strict, "-strict", "-resume", "-anomaly", filepath.Join(strict, "a.tsv"), "-allow", "decode:"+dropReason)); err != nil {
		t.Fatalf("strict allowed: %v", err)
	}
	ids := filepath.Join(t.TempDir(), "ids")
	os.WriteFile(ids, []byte("X1\n"), 0o644)
	file := filepath.Join(t.TempDir(), "f")
	os.WriteFile(file, nil, 0o644)
	for _, bad := range [][]string{
		args(t.TempDir(), "-dump", dump+".none"),
		args(t.TempDir(), "-langreg-root", t.TempDir()),
		args(t.TempDir(), "-ids", ids),
		args(t.TempDir(), "-sort-dir", filepath.Join(file, "s")),
	} {
		if err := run(bad); err == nil {
			t.Errorf("%v: want an error", bad)
		}
	}
	packed := t.TempDir()
	os.MkdirAll(filepath.Join(packed, ".well-known", "geul"), 0o755)
	os.WriteFile(filepath.Join(packed, ".well-known", "geul", "latest.txt"), []byte("x"), 0o644)
	if err := run(args(t.TempDir(), "-langreg-root", packed)); err == nil {
		t.Error("unreadable registry archive: want an error")
	}
}
