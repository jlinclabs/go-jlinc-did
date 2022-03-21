package did

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/nacl/box"
)

type entity struct {
	SigningPublicKey     string
	SigningPrivateKey    string
	EncryptingPublicKey  string
	EncryptingPrivateKey string
	RegistrationSecret   string
}

func createEntity() (entity, error) {
	signPubkey, signPrivkey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return entity{}, err
	}

	encPubkey, encPrivkey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return entity{}, err
	}

	regsecret := make([]byte, 32)
	_, err = rand.Read(regsecret)
	if err != nil {
		return entity{}, err
	}

	return entity{
		SigningPublicKey:     b64Encode(signPubkey),
		SigningPrivateKey:    b64Encode(signPrivkey),
		EncryptingPublicKey:  b64Encode(encPubkey[:]), //[:] to convert *[32]byte to []byte
		EncryptingPrivateKey: b64Encode(encPrivkey[:]),
		RegistrationSecret:   hex.EncodeToString(regsecret)}, nil
}
