//ff:type feature=source type=model
//ff:what dumpSitelink — a sitelink of a dump line: page title and badges
package source

// dumpSitelink is one sitelink of a dump line. The site wikipg stores is the
// map key; Site is decoded but not used.
type dumpSitelink struct {
	Site   string   `json:"site"`
	Title  string   `json:"title"`
	Badges []string `json:"badges"`
}
