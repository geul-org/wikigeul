//ff:func feature=pipeline type=command control=iteration dimension=1
//ff:what checkFlags — check the shared flags after parsing and settle what follows from them; full porter and library commits are required
package pipeline

import (
	"errors"
	"fmt"

	"github.com/geul-org/wikigeul/pkg/port"
)

// checkFlags checks c as BindFlags filled it: -commit and -lib-commit (the
// full git object IDs of the porter and the GEUL library, written in every
// document's TRANSFORMED-BY) are required and must be well formed
// (port.CommitDigest). It settles c.LangVer from langVer and the default of
// LangRoot (-root).
func checkFlags(c *Config, langVer uint) error {
	for _, id := range []string{c.Commit, c.LibCommit} {
		if _, err := port.CommitDigest(id); err != nil {
			return fmt.Errorf("-commit and -lib-commit: %w", err)
		}
	}
	if langVer > 0xFFFFFFFF {
		return errors.New("wikigeul: -langreg-version is a 32-bit version ID")
	}
	c.LangVer = uint32(langVer)
	if c.Workers < 1 || c.Every < 1 {
		return errors.New("wikigeul: -workers and -checkpoint must be at least 1")
	}
	if c.LangRoot == "" {
		c.LangRoot = c.Root
	}
	if c.From == 0 || c.From > c.To {
		return errors.New("wikigeul: need 1 <= -from <= -to")
	}
	return nil
}
