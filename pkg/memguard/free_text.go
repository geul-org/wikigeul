//ff:func feature=memguard type=helper control=sequence
//ff:what freeText — the free share of the tightest limit for log lines, as "12.3%"
package memguard

import "fmt"

// freeText prints Free as a percentage of Size.
func freeText(p pressure) string {
	return fmt.Sprintf("%.1f%%", 100*float64(p.Free)/float64(p.Size))
}
