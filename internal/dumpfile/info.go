//ff:type feature=dumpfile type=model
//ff:what Info — a dump file to convert directly and the provenance its documents record: URL, SHA-256, when it was downloaded
package dumpfile

// Info is a Wikidata JSON dump file read directly (wikigeul -dump): its path
// and file name, size and modification time (Unix seconds; with the name and
// SHA-256 they tell the sort state's dump apart), and the provenance every
// document records (port.Provenance), which wikipg would have put in the
// dump row: the URL, the file's SHA-256 in hex and when it was downloaded
// (Unix seconds).
type Info struct {
	Path         string
	Name         string
	Size         int64
	ModTime      int64
	URL          string
	SHA256       string
	DownloadedAt int64
}
