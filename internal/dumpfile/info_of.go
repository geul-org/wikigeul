//ff:func feature=dumpfile type=loader control=sequence
//ff:what InfoOf — a dump file's Info with wikipg's defaults: URL from the file name's date, SHA-256 from <file>.sha256, download time from the file
package dumpfile

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// InfoOf describes the dump at path as wikipg fills its dump row, so the
// documents record the same provenance as when converted from the loaded
// database (wikipg dumpInfoOf): url, else
// https://dumps.wikimedia.org/wikidatawiki/entities/<date>/<file name> with
// the date of a wikidata-YYYYMMDD-… name; sha256, else the hash for the file
// name in <path>.sha256 (sha256sum format); downloadedAt (Unix seconds), else
// the file's modification time. No URL or no SHA-256 is an error: the
// documents must say which bytes they came from.
func InfoOf(path, url, sha256 string, downloadedAt int64) (Info, error) {
	st, err := os.Stat(path)
	if err != nil {
		return Info{}, err
	}
	name := filepath.Base(path)
	if m := regexp.MustCompile(`wikidata-(\d{8})-`).FindStringSubmatch(name); url == "" && m != nil {
		url = "https://dumps.wikimedia.org/wikidatawiki/entities/" + m[1] + "/" + name
	}
	if url == "" {
		return Info{}, fmt.Errorf("wikigeul: no date in dump name %s: give -dump-url", name)
	}
	if sha256 == "" {
		sha256 = readSum(path+".sha256", name)
	}
	if sha256 == "" {
		return Info{}, fmt.Errorf("wikigeul: no SHA-256 for %s: give -dump-sha256 or write %s.sha256 (sha256sum)", name, path)
	}
	if downloadedAt == 0 {
		downloadedAt = st.ModTime().Unix()
	}
	return Info{Path: path, Name: name, Size: st.Size(), ModTime: st.ModTime().Unix(), URL: url, SHA256: strings.ToLower(sha256), DownloadedAt: downloadedAt}, nil
}
