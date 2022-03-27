package did

const contextUrl = "https://www.w3.org/ns/did/v1"

func RegisterDID(regAddr string) (dataJson string, err error) {
	ent, err := createEntity()
	if err != nil {
		return "", err
	}

	didDoc, sig, err := makeDidDoc(ent)
	if err != nil {
		return "", err
	}

	serverPubkey, err := serverKey(regAddr)
	if err != nil {
		return "", err
	}

	registrantSecret, registrantNonce, err := createRegistrantSecret(serverPubkey, ent)
	if err != nil {
		return "", err
	}

	dataJson, err = register(regAddr, serverPubkey, didDoc, sig, registrantSecret, registrantNonce, ent)
	if err != nil {
		return "", err
	}

	return dataJson, nil
}
