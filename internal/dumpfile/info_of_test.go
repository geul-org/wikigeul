//ff:func feature=dumpfile type=loader control=sequence
//ff:what TestInfoOf — URL from the name's date or given, SHA-256 from the sidecar or given, time from the file or given; missing file, URL or SHA-256 is an error
package dumpfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInfoOf(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "wikidata-20261019-all.json.bz2")
	os.WriteFile(path, []byte("abc"), 0o644)
	os.WriteFile(path+".sha256", []byte("ABCD  wikidata-20261019-all.json.bz2\n"), 0o644)
	in, err := InfoOf(path, "", "", 0)
	st, _ := os.Stat(path)
	if err != nil || in.URL != "https://dumps.wikimedia.org/wikidatawiki/entities/20261019/wikidata-20261019-all.json.bz2" || in.SHA256 != "abcd" ||
		in.DownloadedAt != st.ModTime().Unix() || in.Size != 3 || in.Name != "wikidata-20261019-all.json.bz2" || in.Path != path {
		t.Fatalf("%+v %v", in, err)
	}
	if in, err := InfoOf(path, "https://x", "ff", 7); err != nil || in.URL != "https://x" || in.SHA256 != "ff" || in.DownloadedAt != 7 {
		t.Fatalf("given: %+v %v", in, err)
	}
	other := filepath.Join(dir, "dump.json")
	os.WriteFile(other, nil, 0o644)
	if _, err := InfoOf(other, "", "ff", 0); err == nil {
		t.Fatal("no date, no URL: want an error")
	}
	if _, err := InfoOf(other, "https://x", "", 0); err == nil {
		t.Fatal("no SHA-256: want an error")
	}
	if _, err := InfoOf(filepath.Join(dir, "none"), "https://x", "ff", 0); err == nil {
		t.Fatal("no file: want an error")
	}
}
