//ff:func feature=pipeline type=helper control=sequence
//ff:what Tools — the tools every document records whichever input it came from: wikigeul and the GEUL library at their commits
package pipeline

import "github.com/geul-org/wikigeul/pkg/port"

// Tools is the provenance's tools as geul.org term definition documents
// (GEUL-ORG): wikigeul (document 1, the converter, also when wikipgeul
// export-geul runs it) at commit and the GEUL library (document 2, gogeul) at
// libCommit. The loader is not a tool of the document: docs/provenance.md.
func Tools(commit, libCommit string) []port.Tool {
	return []port.Tool{
		{Doc: 1, Commit: commit},
		{Doc: 2, Commit: libCommit},
	}
}
