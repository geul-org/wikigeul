//ff:func feature=dumpfile type=helper control=selection
//ff:what fixtureSortConfig — a sort of a small dump written to a temporary file (plain, .gz or .bz2 by name), with its SHA-256 and a keep-everything filter
package dumpfile

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/klauspost/compress/gzip"
)

// fixtureSortConfig writes body as name (compressed by its suffix; .bz2
// needs the bzip2 program, else the test is skipped) and returns a sort of
// it into a fresh directory, keeping P and Q, 1 MiB per run.
func fixtureSortConfig(t *testing.T, name, body string) SortConfig {
	t.Helper()
	dir := t.TempDir()
	raw := []byte(body)
	switch {
	case strings.HasSuffix(name, ".gz"):
		var b bytes.Buffer
		zw := gzip.NewWriter(&b)
		zw.Write(raw)
		zw.Close()
		raw = b.Bytes()
	case strings.HasSuffix(name, ".bz2"):
		cmd := exec.Command("bzip2", "-c")
		cmd.Stdin = strings.NewReader(body)
		out, err := cmd.Output()
		if err != nil {
			t.Skipf("bzip2: %v", err)
		}
		raw = out
	}
	path := filepath.Join(dir, name)
	os.WriteFile(path, raw, 0o644)
	sum := sha256.Sum256(raw)
	info, err := InfoOf(path, "https://example.org/d", hex.EncodeToString(sum[:]), 1)
	if err != nil {
		t.Fatal(err)
	}
	return SortConfig{Info: info, Dir: filepath.Join(dir, "sort"), Mem: 1 << 20, Threads: 1, Filter: "test", Log: io.Discard,
		Keep: func(k byte, n uint64) bool { return k == 'P' || k == 'Q' }}
}
