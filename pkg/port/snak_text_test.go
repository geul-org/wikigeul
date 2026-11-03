//ff:func feature=port type=helper control=sequence
//ff:what TestSnakText — property and value words in hex around =; different values differ
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSnakText(t *testing.T) {
	if got := snakText([]core.Word{0x0191, 248}, []core.Word{0x0141, 5}); got != "[0191 00f8]=[0141 0005]" {
		t.Fatalf("%q", got)
	}
	if snakText([]core.Word{1}, []core.Word{2}) == snakText([]core.Word{1}, []core.Word{3}) {
		t.Fatal("values must differ")
	}
}
