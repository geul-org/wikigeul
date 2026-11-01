//ff:func feature=source type=loader control=iteration dimension=1
//ff:what dumpSnaks — the qualifier or reference snaks of a dump line in the dump's property order
package source

// dumpSnaks lists the snaks of m (by property) of statement sid in the order
// wikipgeul numbers their ordinal (orderedSnaks: order, then properties it
// misses, sorted), each in qualifier form (dumpQualifier).
func dumpSnaks(sid int64, order []string, m map[string][]dumpSnak) ([]Qualifier, error) {
	var out []Qualifier
	for _, s := range orderedSnaks(order, m) {
		q, err := dumpQualifier(sid, s)
		if err != nil {
			return nil, err
		}
		out = append(out, q)
	}
	return out, nil
}
