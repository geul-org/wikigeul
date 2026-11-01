//ff:func feature=pipeline type=helper control=iteration dimension=1
//ff:what WriteReport — the run report as TSV: a # line naming the run, a header, one line per entity
package pipeline

import (
	"os"
	"strings"

	"github.com/geul-org/wikigeul/pkg/port"
)

// WriteReport writes reps to path as TSV (reportLine) after the comment line
// head (RunLine) and a header line.
func WriteReport(path, head string, reps []port.Report) error {
	lines := []string{head, reportHeader}
	for _, r := range reps {
		lines = append(lines, reportLine(r))
	}
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0o644)
}
