//ff:func feature=pipeline type=helper control=sequence
//ff:what TestWriteReport — the run comment, a header and one line per report; the header (id, status, …) has as many columns as a row, a missing entity's row too
package pipeline

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/wikigeul/pkg/port"
)

func TestWriteReport(t *testing.T) {
	p := filepath.Join(t.TempDir(), "r.tsv")
	if err := WriteReport(p, "# run", []port.Report{{ID: "Q1"}, {ID: "Q2", Missing: true}}); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(p)
	if lines := strings.Split(string(b), "\n"); len(lines) != 5 || lines[0] != "# run" || !strings.HasPrefix(lines[1], "id\tstatus\t") || !strings.HasPrefix(lines[3], "Q2\tmissing\t") ||
		strings.Count(lines[1], "\t") != strings.Count(lines[2], "\t") || strings.Count(lines[1], "\t") != strings.Count(lines[3], "\t") {
		t.Fatalf("got %q", b)
	}
}
