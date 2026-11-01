//ff:func feature=source type=helper control=sequence
//ff:what numericText — a JSON number literal as PostgreSQL numeric prints it after jsonb stored it: plain decimal, scale from the literal
package source

import (
	"strconv"
	"strings"
)

// numericText prints the JSON number lit as numeric_out prints numeric_in's
// value (jsonb keeps numbers as numeric): no exponent; the display scale is
// the literal's fraction digits minus its exponent, at least 0, so trailing
// zeros of the literal stay (1.50 → 1.50, 1.0e-5 → 0.000010, 15e2 → 1500);
// leading zeros go; zero has no sign (-0.0 → 0.0).
func numericText(lit string) string {
	neg := strings.HasPrefix(lit, "-")
	s := strings.TrimPrefix(lit, "-")
	exp := 0
	if i := strings.IndexAny(s, "eE"); i >= 0 {
		exp, _ = strconv.Atoi(s[i+1:])
		s = s[:i]
	}
	whole, frac, _ := strings.Cut(s, ".")
	digits := whole + frac
	point := len(whole) + exp // digits before the decimal point
	ip := "0"
	if point > 0 {
		ip = (digits + strings.Repeat("0", max(0, point-len(digits))))[:point]
		ip = strings.TrimLeft(ip, "0")
	}
	if ip == "" {
		ip = "0"
	}
	if strings.Trim(digits, "0") == "" {
		neg = false
	}
	out := ip
	if scale := len(frac) - exp; scale > 0 {
		out += "." + strings.Repeat("0", max(0, -point)) + digits[max(0, point):]
	}
	if neg {
		return "-" + out
	}
	return out
}
