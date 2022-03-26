package did

import (
	"crypto/rand"
	"io"

	"golang.org/x/crypto/nacl/box"
)

func createRegistrantSecret(serverKey string, ent entity) (string, string, error) {
	registerSecret := ent.RegistrationSecret

	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return "", "", err
	}

	var out []byte
	registrantSecret := box.Seal(out, []byte(registerSecret), &nonce, (*[32]byte)(b64Decode(serverKey)), (*[32]byte)(b64Decode(ent.EncryptingPrivateKey)))

	return b64Encode(registrantSecret), b64Encode(nonce[:]), nil
}
