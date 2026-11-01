//ff:func feature=source type=loader control=sequence
//ff:what dumpValueEntity — the entity ID a dump datavalue points at, as wikipg fills value_entity; empty for other values
package source

import "encoding/json"

// dumpValueEntity returns the "id" of a wikibase-entityid datavalue and ""
// for any other type, a value that is not an object or has no id (wikipg
// internal/dump.ValueEntity; "" is stored as NULL).
func dumpValueEntity(dv *dumpDataValue) string {
	if dv.Type != "wikibase-entityid" {
		return ""
	}
	var v struct {
		ID string `json:"id"`
	}
	if json.Unmarshal(dv.Value, &v) != nil {
		return ""
	}
	return v.ID
}
