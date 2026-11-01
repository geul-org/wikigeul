//ff:type feature=source type=model
//ff:what dumpEntity — one Wikidata JSON dump line as wikipgeul decodes it: ID, datatype, revision, terms, claims, sitelinks
package source

// dumpEntity is one entity line of the Wikidata JSON dump decoded the way
// wikipgeul decodes it (wikipgeul internal/dump.Entity, encoding/json into maps):
// the fields the loaded tables keep for an item or property. Unknown fields
// are ignored as wikipgeul ignores them; the fields it reads but the tables do
// not need (type, statement GUID, snak hash, sitelink site) are decoded too,
// so a line that fails wikipgeul's decoding fails here as well (Lexeme fields
// are left out: an item or property line has none). A JSON object key given
// twice keeps the last value (encoding/json), as in wikipgeul.
type dumpEntity struct {
	ID           string                     `json:"id"`
	Type         string                     `json:"type"`
	Datatype     string                     `json:"datatype"`
	LastRevID    int64                      `json:"lastrevid"`
	Modified     string                     `json:"modified"`
	Labels       map[string]dumpTerm        `json:"labels"`
	Descriptions map[string]dumpTerm        `json:"descriptions"`
	Aliases      map[string][]dumpTerm      `json:"aliases"`
	Claims       map[string][]dumpStatement `json:"claims"`
	Sitelinks    map[string]dumpSitelink    `json:"sitelinks"`
}
