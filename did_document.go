package did

import (
	"crypto/ed25519"
	"fmt"
	"time"
)

const ISOStringMillisec = "2006-01-02T15:04:05.999Z"

type didDoc struct {
	AtContext  string     `json:"@context"`
	ID         string     `json:"id"`
	CreatedAt  string     `json:"created"`
	PublicKeys publicKeys `json:"publicKey"`
}

type publicKeys []pubkey

type pubkey struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	Controller      string `json:"controller"`
	PublicKeyBase64 string `json:"publicKeyBase64"`
	PublicKeyBase58 string `json:"publicKeyBase58"`
}

func makeDidDoc(ent entity) (doc didDoc, sig string, err error) {
	didId := "did:jlinc:" + ent.SigningPublicKey
	created := fmt.Sprintf("%s", time.Now().UTC().Format(ISOStringMillisec))

	didDocument := didDoc{
		AtContext: contextUrl,
		ID:        didId,
		CreatedAt: created,
	}
	signing := pubkey{
		ID:              didId + "#signing",
		Type:            "Ed25519VerificationKey2018",
		Controller:      didId,
		PublicKeyBase64: ent.SigningPublicKey,
		PublicKeyBase58: b64tob58(ent.SigningPublicKey),
	}
	encrypting := pubkey{
		ID:              didId + "#encrypting",
		Type:            "X25519KeyAgreementKey2019",
		Controller:      didId,
		PublicKeyBase64: ent.EncryptingPublicKey,
		PublicKeyBase58: b64tob58(ent.EncryptingPublicKey),
	}
	var pubkeys publicKeys
	pubkeys = append(pubkeys, signing)
	pubkeys = append(pubkeys, encrypting)
	didDocument.PublicKeys = pubkeys

	toBeSigned := didId + "." + created
	toBeSignedHashed := getHash(toBeSigned)
	signer := b64Decode(ent.SigningPrivateKey)
	signature := ed25519.Sign(signer, toBeSignedHashed)
	sig = b64Encode(signature)

	return didDocument, sig, nil
}
