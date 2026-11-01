//ff:func feature=source type=loader control=iteration dimension=2
//ff:what dumpClaims — the claims of a dump line in statement ID order: property ID as text, then dump order
package source

import (
	"maps"
	"slices"
)

// dumpClaims lists m (property → statements) in the order wikipgeul numbers
// statement_id (properties sorted as text, each property's statements in
// array order) and numbers them from 0 in that order: the number stands for
// the statement_id, which the documents never show. The property of each
// claim is its main snak's, as wikipgeul stores it.
func dumpClaims(m map[string][]dumpStatement) ([]Claim, error) {
	var out []Claim
	for _, p := range slices.Sorted(maps.Keys(m)) {
		for _, s := range m[p] {
			c, err := dumpClaim(int64(len(out)), s)
			if err != nil {
				return nil, err
			}
			out = append(out, c)
		}
	}
	return out, nil
}
