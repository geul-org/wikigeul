//ff:func feature=source type=loader control=sequence
//ff:what dumpClaim — one claim of a dump line with its qualifiers and references, numbered id
package source

// dumpClaim makes statement s, numbered id, a Claim: the main snak's
// property, value type, datatype, value and snak type (dumpQualifier),
// the rank, the qualifiers in qualifiers-order (dumpSnaks) and the
// references in order (dumpReferences).
func dumpClaim(id int64, s dumpStatement) (Claim, error) {
	m, err := dumpQualifier(id, s.Mainsnak)
	if err != nil {
		return Claim{}, err
	}
	qs, err := dumpSnaks(id, s.QualifiersOrder, s.Qualifiers)
	if err != nil {
		return Claim{}, err
	}
	rs, err := dumpReferences(id, s.References)
	if err != nil {
		return Claim{}, err
	}
	return Claim{ID: id, Property: m.Property, Type: m.Type, Datatype: m.Datatype, Value: m.Value, Rank: s.Rank, SnakType: m.SnakType,
		Qualifiers: qs, References: rs}, nil
}
