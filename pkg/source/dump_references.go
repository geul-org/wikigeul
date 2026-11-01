//ff:func feature=source type=loader control=iteration dimension=1
//ff:what dumpReferences — the references of a claim of a dump line in order, each with this occurrence's snaks in snaks-order
package source

// dumpReferences lists rs, the references of statement sid, in array order
// (ordinal) with their hash and this occurrence's own snaks in snaks-order
// (dumpSnaks), as the PG query reads them with reference_snak_order. A
// reference with no snaks is left out: the PG query joins statement_reference
// to reference_snak and so returns no row for it.
func dumpReferences(sid int64, rs []dumpReference) ([]Reference, error) {
	var out []Reference
	for ord, r := range rs {
		snaks, err := dumpSnaks(sid, r.SnaksOrder, r.Snaks)
		if err != nil {
			return nil, err
		}
		if len(snaks) > 0 {
			out = append(out, Reference{StatementID: sid, Ordinal: ord, Hash: r.Hash, Snaks: snaks})
		}
	}
	return out, nil
}
