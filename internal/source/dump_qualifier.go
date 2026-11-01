//ff:func feature=source type=loader control=sequence
//ff:what dumpQualifier — one snak of a dump line in a loaded qualifier's form: statement, property, value type, datatype, value, snak type
package source

// dumpQualifier makes snak s of statement sid a Qualifier, the form of a
// qualifier or reference snak row (and of a claim's main snak): property,
// value type and value (dumpSnakValue), datatype and snak type.
func dumpQualifier(sid int64, s dumpSnak) (Qualifier, error) {
	typ, v, err := dumpSnakValue(s.DataValue)
	if err != nil {
		return Qualifier{}, err
	}
	return Qualifier{StatementID: sid, Property: s.Property, Type: typ, Datatype: s.Datatype, Value: v, SnakType: s.SnakType}, nil
}
