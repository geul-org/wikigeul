//ff:type feature=port type=model
//ff:what codes — the GEUL codes this porter writes, named once (the library's where it has them), and the Wikidata constants it reads
package port

import "github.com/geul-org/gogeul/pkg/core"

// The GEUL codes this porter writes. The Value and Meta codes and the Edge
// codes the library names come from core, so no package keeps its own copy of
// a number; edge/REFERENCE-NODE, LABEL, ALIAS and DESCRIPTION, which core does
// not name, are their Edge Codec (code - 0xA0, TAG 0x0A00 + Codec<<4 + L).
const (
	codecURIEven              = core.CodecURIEven           // Value URI, even bytes (0x074L; odd 0x075L)
	codecLanguageShort        = core.CodecLanguageShortCode // Value LANGUAGE-SHORT-CODE (0x07AL)
	edgeReferenceNode  uint16 = 0x00                        // Edge REFERENCE-NODE (0x0A0L)
	edgeLabel          uint16 = 0x04                        // Edge LABEL (0x0A4L)
	edgeAlias          uint16 = 0x05                        // Edge ALIAS (0x0A5L)
	edgeDescription    uint16 = 0x06                        // Edge DESCRIPTION (0x0A6L)
	edgeTriple                = core.EdgeTriple             // Edge TRIPLE (0x0AAL)
	edgeQualifier             = core.EdgeQualifier          // Edge QUALIFIER (0x0ABL)
	edgeSource                = core.EdgeSource             // Edge SOURCE (0x0C0L, 근거)
	metaTransformedBy         = core.MetaTransformedBy      // Meta TRANSFORMED-BY (0x1D4L)

	schemaOrg = "http://schema.org/" // schema.org namespace as Wikidata's RDF writes it (sitelink schema:about)

	gregorian = "http://www.wikidata.org/entity/Q1985727" // Wikidata calendar model: proleptic Gregorian
	julian    = "http://www.wikidata.org/entity/Q1985786" // Wikidata calendar model: proleptic Julian
)
