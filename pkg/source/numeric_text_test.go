//ff:func feature=source type=helper control=iteration dimension=1
//ff:what TestNumericText — JSON number literals print as PostgreSQL numeric: no exponent, scale from the literal, no sign on zero
package source

import "testing"

func TestNumericText(t *testing.T) {
	for _, c := range []struct{ in, want string }{
		{"11", "11"}, {"0", "0"}, {"-0", "0"}, {"-0.0", "0.0"}, {"1.50", "1.50"}, {"-7.25", "-7.25"},
		{"0.00027777777777778", "0.00027777777777778"}, {"1.0e-5", "0.000010"}, {"1e-5", "0.00001"},
		{"2.5E-3", "0.0025"}, {"15e2", "1500"}, {"1.5e1", "15"}, {"1.55e1", "15.5"}, {"1E+2", "100"}, {"007", "7"},
	} {
		if got := numericText(c.in); got != c.want {
			t.Errorf("%s: got %s, want %s", c.in, got, c.want)
		}
	}
}
