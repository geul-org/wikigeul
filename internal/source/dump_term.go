//ff:type feature=source type=model
//ff:what dumpTerm — a label, description or alias of a dump line: language and text
package source

// dumpTerm is a label, description or alias value of a dump line. The
// language wikipg stores is the map key, not Language.
type dumpTerm struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}
