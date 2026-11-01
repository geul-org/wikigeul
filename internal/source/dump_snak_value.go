//ff:func feature=source type=loader control=selection
//ff:what dumpSnakValue — a dump snak's value type and value in the form the PG queries give it
package source

import "bytes"

// dumpSnakValue gives the value type and value of datavalue dv as the PG
// queries read the columns wikipg wrote from it (wikipg snakValues):
//   - no datavalue (novalue, somevalue): both empty;
//   - the value JSON contains the escape \u0000 (jsonb cannot hold it, so
//     wikipg keeps it as text in value_json): exactValue of that text — a
//     JSON string decoded, anything else as written in the dump;
//   - an entity value with an id: the bare ID (value_entity);
//   - otherwise value #>> '{}' (jsonbText): a string unquoted, anything else
//     as PostgreSQL prints the jsonb.
func dumpSnakValue(dv *dumpDataValue) (string, string, error) {
	switch {
	case dv == nil:
		return "", "", nil
	case bytes.Contains(dv.Value, []byte(`\u0000`)):
		return dv.Type, exactValue("", string(dv.Value)), nil
	case dumpValueEntity(dv) != "":
		return dv.Type, dumpValueEntity(dv), nil
	default:
		v, err := jsonbText(dv.Value)
		return dv.Type, v, err
	}
}
