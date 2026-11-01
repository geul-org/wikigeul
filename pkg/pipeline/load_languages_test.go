//ff:func feature=pipeline type=loader control=sequence
//ff:what TestLoadLanguages — the registry at the root goes into the config; a root without one and an unreadable pack archive fail
package pipeline

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadLanguages(t *testing.T) {
	c := Config{LangRoot: t.TempDir()}
	fixtureRegistry(t, c.LangRoot)
	if err := LoadLanguages(&c); err != nil || len(c.Langs) != 3 || c.LangDoc == "" {
		t.Fatalf("%v %v %q", err, c.Langs, c.LangDoc)
	}
	if err := LoadLanguages(&Config{LangRoot: t.TempDir()}); err == nil {
		t.Fatal("no registry: want error")
	}
	packed := t.TempDir()
	os.MkdirAll(filepath.Join(packed, ".well-known", "geul"), 0o755)
	os.WriteFile(filepath.Join(packed, ".well-known", "geul", "latest.txt"), []byte("x"), 0o644)
	if err := LoadLanguages(&Config{LangRoot: packed}); err == nil {
		t.Fatal("unreadable archive: want error")
	}
}
