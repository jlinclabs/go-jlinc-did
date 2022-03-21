package did

import "log"

type PrivateDidData struct {
	SigningPublicKey     string
	SigningPrivateKey    string
	EncryptingPublicKey  string
	EncryptingPrivateKey string
	RegistrationSecret   string
	DID                  string
}

func RegisterDID(regAddr string) (string, error) {
	entity, err := createEntity()
	if err != nil {
		return "", err
	}

	did, err := makeDidDoc(entity)
	if err != nil {
		return "", err
	}

	dataJson, err := register(did)
	if err != nil {
		return "", err
	}

	log.Printf("entity: %v", entity)
	return dataJson, nil
}
