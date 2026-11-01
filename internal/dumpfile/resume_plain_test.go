//ff:func feature=dumpfile type=loader control=sequence
//ff:what TestResumePlain — offset 0 is untouched; a saved offset seeks there and restores the hash so the whole file's SHA-256 comes out; a bad hash state is an error
package dumpfile

import (
	"crypto/sha256"
	"encoding"
	"encoding/hex"
	"io"
	"testing"
)

func TestResumePlain(t *testing.T) {
	const body = "[\n{\"id\":\"Q1\"}\n]\n"
	c := fixtureSortConfig(t, "d.json", body)
	h := sha256.New()
	h.Write([]byte(body[:2]))
	state, _ := h.(encoding.BinaryMarshaler).MarshalBinary()
	in, err := openInput(c, State{Offset: 2, Hash: state})
	if err != nil {
		t.Fatal(err)
	}
	rest, _ := io.ReadAll(in.r)
	in.h.Write(rest)
	sum, _ := in.finish()
	in.close()
	want := sha256.Sum256([]byte(body))
	if string(rest) != body[2:] || sum != hex.EncodeToString(want[:]) {
		t.Fatalf("%q %s", rest, sum)
	}
	in, _ = openInput(c, State{})
	if err := resumePlain(in, State{}); err != nil {
		t.Fatal(err)
	}
	if err := resumePlain(in, State{Offset: 2, Hash: []byte("bad")}); err == nil {
		t.Fatal("bad hash state: want an error")
	}
	in.close()
}
