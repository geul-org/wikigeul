//ff:type feature=source type=model
//ff:what dumpReference — one reference of a claim of a dump line: hash, snaks by property, snaks-order
package source

// dumpReference is one reference of a claim: its hash and its snaks by
// property with snaks-order (wikipgeul internal/dump.Reference).
type dumpReference struct {
	Hash       string                `json:"hash"`
	Snaks      map[string][]dumpSnak `json:"snaks"`
	SnaksOrder []string              `json:"snaks-order"`
}
