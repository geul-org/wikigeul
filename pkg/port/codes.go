//ff:type feature=port type=model
//ff:what codes — the GEUL codes this porter writes, named once (the library's where it has them), and the Wikidata constants it reads
package port

import "github.com/geul-org/gogeul/pkg/core"

// The GEUL codes this porter writes. The Value and Meta codes and the Edge
// codes the library names come from core, so no package keeps its own copy of
// a number; edge/REFERENCE-NODE, LABEL, ALIAS and DESCRIPTION, which core does
// not name, are their Edge Codec (code - 0x80, TAG 0x0800 + Codec<<4 + L).
const (
	codecURIEven              = core.CodecURIEven           // Value URI, even bytes (0x054L; odd 0x055L)
	codecLanguageShort        = core.CodecLanguageShortCode // Value LANGUAGE-SHORT-CODE (0x05AL)
	edgeReferenceNode  uint16 = 0x00                        // Edge REFERENCE-NODE (0x080L)
	edgeLabel          uint16 = 0x04                        // Edge LABEL (0x084L)
	edgeAlias          uint16 = 0x05                        // Edge ALIAS (0x085L)
	edgeDescription    uint16 = 0x06                        // Edge DESCRIPTION (0x086L)
	edgeTriple                = core.EdgeTriple             // Edge TRIPLE (0x08AL)
	edgeQualifier             = core.EdgeQualifier          // Edge QUALIFIER (0x08BL)
	edgeSource                = core.EdgeSource             // Edge SOURCE (0x0A0L, 근거)
	metaTransformedBy         = core.MetaTransformedBy      // Meta TRANSFORMED-BY (0x1D4L)

	schemaOrg = "http://schema.org/" // schema.org namespace as Wikidata's RDF writes it (sitelink schema:about)

	gregorian = "http://www.wikidata.org/entity/Q1985727" // Wikidata calendar model: proleptic Gregorian
	julian    = "http://www.wikidata.org/entity/Q1985786" // Wikidata calendar model: proleptic Julian
)
