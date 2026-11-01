//ff:func feature=source type=loader control=sequence
//ff:what DecodeDumpEntity — one Wikidata JSON dump line as the Entity wikipgeul export-geul reads after wikipgeul loaded that line
package source

import (
	"encoding/json"
	"errors"
)

// DecodeDumpEntity decodes raw, the JSON object of one dump line (without
// the separating comma), into the Entity that wikipgeul pgsource.LoadEntity returns for the same
// entity once wikipgeul has loaded the line (docs/dump.md 대응표): head
// (datatype, lastrevid, modified as POSIX seconds), labels and descriptions
// by language code in byte order, aliases by language then dump order,
// claims by property ID as text then dump order (dumpClaims), each with its
// qualifiers (qualifiers-order) and references (each occurrence's own
// snaks-order), and sitelinks by site ID with their badges. Values take the
// form the PG queries give (dumpSnakValue: a bare entity ID, an unquoted
// string, else the value as PostgreSQL jsonb prints it). A line wikipgeul cannot
// decode (encoding/json into the same field types) or without an ID is an
// error; wikipgeul records such a line as a load_error and loads no entity.
func DecodeDumpEntity(raw []byte) (Entity, error) {
	var d dumpEntity
	if err := json.Unmarshal(raw, &d); err != nil {
		return Entity{}, err
	}
	if d.ID == "" {
		return Entity{}, errors.New("wikigeul: dump entity without id")
	}
	modified, err := dumpModified(d.Modified)
	if err != nil {
		return Entity{}, err
	}
	claims, err := dumpClaims(d.Claims)
	if err != nil {
		return Entity{}, err
	}
	return Entity{
		ID:           d.ID,
		Datatype:     d.Datatype,
		LastRevID:    d.LastRevID,
		Modified:     modified,
		Labels:       dumpTexts(d.Labels),
		Descriptions: dumpTexts(d.Descriptions),
		Aliases:      dumpAliases(d.Aliases),
		Claims:       claims,
		Sitelinks:    dumpSitelinks(d.Sitelinks),
	}, nil
}
