//ff:func feature=pipeline type=command control=sequence
//ff:what packTestProv — a provenance with the config's commits
package pipeline

import "github.com/geul-org/wikigeul/pkg/port"

// packTestProv is a dump's provenance with c's tools.
func packTestProv(c Config) port.Provenance {
	return port.Provenance{DumpURI: "u", Tools: Tools(c.Commit, c.LibCommit)}
}
