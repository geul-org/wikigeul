//ff:func feature=pipeline type=helper control=sequence
//ff:what RunLine — the report's first line: the run's range, dump and commits as a # comment
package pipeline

import "fmt"

// RunLine is the comment that opens the report: the QID range, the loaded
// dump (schema, dump_id), the porter and library commits and the GEUL Language
// Registry version and content hash (LangDoc) of run c. The documents depend
// on the dump, the commits and that registry version only; the range says
// which documents this report covers. A dump read directly (-dump) is named
// by c.Source (file name and SHA-256) instead of schema and dump_id.
func RunLine(c Config) string {
	if c.Source != "" {
		return fmt.Sprintf("# wikigeul -from %d -to %d -dump %s -commit %s -lib-commit %s -langreg %s", c.From, c.To, c.Source, c.Commit, c.LibCommit, c.LangDoc)
	}
	return fmt.Sprintf("# wikigeul -from %d -to %d -schema %s -dump-id %d -commit %s -lib-commit %s -langreg %s", c.From, c.To, c.Schema, c.Dump, c.Commit, c.LibCommit, c.LangDoc)
}
