//ff:type feature=source type=model
//ff:what dumpSnak — a main, qualifier or reference snak of a dump line: snak type, property, datatype, datavalue
package source

// dumpSnak is a snak of a dump line (wikipg internal/dump.Snak; the snak
// hash is decoded but not used).
type dumpSnak struct {
	SnakType  string         `json:"snaktype"`
	Property  string         `json:"property"`
	Hash      string         `json:"hash"`
	Datatype  string         `json:"datatype"`
	DataValue *dumpDataValue `json:"datavalue"`
}
