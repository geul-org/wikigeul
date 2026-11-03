//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestIsQIDTagSpace — over all 512 codes only WIKIDATA-QID × 5Q (0x014L~0x018L) is a QID; markers, pointers, other schemes, ISSUER codes and every other Kind are not
package port

import "testing"

func TestIsQIDTagSpace(t *testing.T) {
	n := 0
	for _, c := range tagSpaceCodes() {
		want := c.Code >= 0x014 && c.Code <= 0x018 // §3.1 WIKIDATA-QID Identified~Variable
		r := tagRecord(c)
		if got := isQID(r); got != want {
			t.Errorf("0x%03XL: isQID %v want %v", c.Code, got, want)
		}
		if want {
			n++
		}
	}
	if n != 5 {
		t.Fatalf("%d QID codes, want 5", n)
	}
}
