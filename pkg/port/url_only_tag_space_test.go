//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestURLOnlyTagSpace — over all 512 codes only URI even and odd (0x074L, 0x075L) as the one P854 value is URL-only
package port

import "testing"

func TestURLOnlyTagSpace(t *testing.T) {
	for _, c := range tagSpaceCodes() {
		want := c.Code == 0x074 || c.Code == 0x075 // §3.4 URI-EVEN, URI-ODD
		if got := urlOnly([]refSnak{{Property: "P854", Value: tagRecord(c)}}); got != want {
			t.Errorf("0x%03XL: urlOnly %v want %v", c.Code, got, want)
		}
	}
}
