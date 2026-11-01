//ff:func feature=source type=loader control=iteration dimension=1
//ff:what dumpSitelinks — the sitelinks of a dump line by site ID in byte order, badges in dump order
package source

import (
	"maps"
	"slices"
)

// dumpSitelinks lists m (site → sitelink) as the PG query reads sitelink:
// the map key as the site, by its bytes, the title exact and the badges in
// dump order (an empty list when there are none, as the '{}' array reads).
func dumpSitelinks(m map[string]dumpSitelink) []Sitelink {
	var out []Sitelink
	for _, site := range slices.Sorted(maps.Keys(m)) {
		badges := m[site].Badges
		if badges == nil {
			badges = []string{}
		}
		out = append(out, Sitelink{Site: site, Title: m[site].Title, Badges: badges})
	}
	return out
}
