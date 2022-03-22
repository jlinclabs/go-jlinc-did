package did

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"

	"github.com/shengdoushi/base58"
)

func getHash(j string) []byte {
	h := sha256.New()
	h.Write([]byte(j))
	return h.Sum(nil)
}

func getByteHash(j []byte) []byte {
	h := sha256.New()
	h.Write(j)
	return h.Sum(nil)
}

func b64Decode(s string) []byte {
	decoded, _ := base64.RawURLEncoding.DecodeString(s)
	return decoded
}

func b64Encode(h []byte) string {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.RawURLEncoding, &buf)
	encoder.Write(h)
	encoder.Close()
	return buf.String()
}

func b58Decode(s string) []byte {
	decoded, _ := base58.Decode(s, base58.BitcoinAlphabet)
	return decoded
}

func b58Encode(h []byte) string {
	return base58.Encode(h, base58.BitcoinAlphabet)
}

func b58tob64(s string) string {
	return b64Encode(b58Decode(s))
}

func b64tob58(s string) string {
	return b58Encode(b64Decode(s))
}
