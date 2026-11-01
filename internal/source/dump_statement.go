//ff:type feature=source type=model
//ff:what dumpStatement — one claim of a dump line: rank, main snak, qualifiers with their order, references
package source

// dumpStatement is one claim of a dump line: rank, main snak, qualifiers by
// property with qualifiers-order and references in order (wikipg
// internal/dump.Statement; the GUID is decoded but not used).
type dumpStatement struct {
	ID              string                `json:"id"`
	Rank            string                `json:"rank"`
	Mainsnak        dumpSnak              `json:"mainsnak"`
	Qualifiers      map[string][]dumpSnak `json:"qualifiers"`
	QualifiersOrder []string              `json:"qualifiers-order"`
	References      []dumpReference       `json:"references"`
}
