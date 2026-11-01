//ff:func feature=pipeline type=command control=iteration dimension=1
//ff:what TestBindFlags — defaults and overrides (registry root and version too); a missing or short commit, a bad range, an unknown flag, a registry version over 32 bits and zero workers or checkpoint fail
package pipeline

import (
	"flag"
	"strings"
	"testing"
)

func TestBindFlags(t *testing.T) {
	a, b := strings.Repeat("a", 40), strings.Repeat("b", 40)
	parse := func(args ...string) (Config, error) {
		fs := flag.NewFlagSet("t", flag.ContinueOnError)
		var c Config
		check := BindFlags(fs, &c, 7, "workers")
		if err := fs.Parse(args); err != nil {
			return c, err
		}
		return c, check()
	}
	c, err := parse("-commit", a, "-lib-commit", b)
	if err != nil || c.Root != "." || c.From != 1 || c.To != 1000 || c.Commit != a || c.LibCommit != b || c.LangRoot != "." || c.LangVer != 0 || c.Workers != 7 || c.Every != 100000 || c.Samples != 2048 || c.Strict || c.Anomaly != "wikigeul-anomaly.tsv" || c.Allow != nil {
		t.Fatalf("defaults: %+v %v", c, err)
	}
	c, err = parse("-commit", a, "-lib-commit", b, "-from", "5", "-to", "9", "-report", "r.tsv", "-root", "/x", "-langreg-root", "/y", "-langreg-version", "2")
	if err != nil || c.From != 5 || c.To != 9 || c.Report != "r.tsv" || c.Root != "/x" || c.LangRoot != "/y" || c.LangVer != 2 {
		t.Fatalf("overrides: %+v %v", c, err)
	}
	c, err = parse("-commit", a, "-lib-commit", b, "-strict", "-anomaly", "x.tsv", "-allow", "raw:time: 0년", "-allow", "warn:E7")
	if err != nil || !c.Strict || c.Anomaly != "x.tsv" || len(c.Allow) != 2 || c.Allow[1] != "warn:E7" {
		t.Fatalf("strict: %+v %v", c, err)
	}
	for _, args := range [][]string{{}, {"-commit", a}, {"-commit", "abc", "-lib-commit", b}, {"-commit", a, "-lib-commit", b, "-from", "0"}, {"-commit", a, "-lib-commit", b, "-from", "9", "-to", "5"}, {"-nope"}, {"-commit", a, "-lib-commit", b, "-langreg-version", "4294967296"}, {"-commit", a, "-lib-commit", b, "-workers", "0"}, {"-commit", a, "-lib-commit", b, "-checkpoint", "0"}} {
		if _, err := parse(args...); err == nil {
			t.Fatalf("%v: want error", args)
		}
	}
}
