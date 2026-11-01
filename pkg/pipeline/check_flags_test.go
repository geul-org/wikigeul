//ff:func feature=pipeline type=command control=iteration dimension=1
//ff:what TestCheckFlags — settles the registry version and root; a short commit, a version over 32 bits, zero workers or checkpoint and a bad range fail
package pipeline

import (
	"strings"
	"testing"
)

func TestCheckFlags(t *testing.T) {
	a := strings.Repeat("a", 40)
	good := Config{Root: "/x", From: 1, To: 2, Commit: a, LibCommit: a, Workers: 1, Every: 1}
	c := good
	if err := checkFlags(&c, 2); err != nil || c.LangVer != 2 || c.LangRoot != "/x" {
		t.Fatalf("%+v %v", c, err)
	}
	c = good
	c.LangRoot = "/y"
	if err := checkFlags(&c, 0); err != nil || c.LangRoot != "/y" {
		t.Fatalf("given root: %+v %v", c, err)
	}
	for name, bad := range map[string]func(*Config){
		"commit":     func(c *Config) { c.Commit = "abc" },
		"lib commit": func(c *Config) { c.LibCommit = "" },
		"workers":    func(c *Config) { c.Workers = 0 },
		"checkpoint": func(c *Config) { c.Every = 0 },
		"from":       func(c *Config) { c.From = 0 },
		"range":      func(c *Config) { c.From = 9 },
	} {
		c := good
		bad(&c)
		if err := checkFlags(&c, 0); err == nil {
			t.Errorf("%s: want error", name)
		}
	}
	c = good
	if err := checkFlags(&c, 1<<32); err == nil {
		t.Fatal("version over 32 bits: want error")
	}
}
