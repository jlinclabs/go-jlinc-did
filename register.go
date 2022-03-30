package did

import (
	"bytes"
	"crypto/ed25519"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type PrivateDidData struct {
	SigningPublicKey     string `json:"signingPublicKey"`
	SigningPrivateKey    string `json:"signingPrivateKey"`
	EncryptingPublicKey  string `json:"encryptingPublicKey"`
	EncryptingPrivateKey string `json:"encryptingPrivateKey"`
	RegistrationSecret   string `json:"registrationSecret"`
	DID                  string `json:"did"`
}
type regRequest struct {
	DID       didDoc    `json:"did"`
	Signature string    `json:"signature"`
	Secret    regSecret `json:"secret"`
}
type regSecret struct {
	Cyphertext string `json:"cyphertext"`
	Nonce      string `json:"nonce"`
}
type regResponse struct {
	ID        string `json:"id"`
	Challenge string `json:"challenge"`
}
type confirmRequest struct {
	ChallengeResponse string `json:"challengeResponse"`
}
type confirmResponse struct {
	DID string `json:"id"`
}

func register(regAddr string, serverPubkey string, didDoc didDoc, sig string, registrantSecret string, registrantNonce string, ent entity) (string, error) {
	regSec := regSecret{
		Cyphertext: registrantSecret,
		Nonce:      registrantNonce,
	}
	regReq := regRequest{
		DID:       didDoc,
		Signature: sig,
		Secret:    regSec,
	}

	// make register request
	jsonData, err := json.Marshal(regReq)
	if err != nil {
		return "", err
	}
	requestAddr := regAddr + "/register"
	resp, err := http.Post(requestAddr, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	j, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()

	// make registration confirmation
	var rr regResponse
	err = json.Unmarshal(j, &rr)
	if err != nil {
		return "", err
	}

	signer := b64Decode(ent.SigningPrivateKey)
	signature := ed25519.Sign(signer, getHash(rr.Challenge))
	sig = b64Encode(signature)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        rr.ID,
		"signature": sig,
	})
	tokenString, err := token.SignedString([]byte(ent.RegistrationSecret))
	if err != nil {
		return "", err
	}
	confirm := confirmRequest{
		ChallengeResponse: tokenString,
	}
	confirmMessage, err := json.Marshal(confirm)
	if err != nil {
		return "", err
	}
	confirmAddr := regAddr + "/confirm"
	confirmResp, err := http.Post(confirmAddr, "text/plain", bytes.NewBuffer(confirmMessage))
	if err != nil {
		return "", err
	}
	c, err := ioutil.ReadAll(confirmResp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	var cr confirmResponse
	err = json.Unmarshal(c, &cr)
	if err != nil {
		return "", err
	}

	dataToSave := PrivateDidData{
		SigningPublicKey:     ent.SigningPublicKey,
		SigningPrivateKey:    ent.SigningPrivateKey,
		EncryptingPublicKey:  ent.EncryptingPublicKey,
		EncryptingPrivateKey: ent.EncryptingPrivateKey,
		RegistrationSecret:   ent.RegistrationSecret,
		DID:                  cr.DID,
	}
	jsonToSave, err := json.Marshal(dataToSave)
	if err != nil {
		return "", err
	}

	return string(jsonToSave), nil
}
