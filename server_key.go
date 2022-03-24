package did

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ServerMasterKey struct {
	MasterPublicKey string `json:"masterPublicKey"`
}

func serverKey(regAddr string) (string, error) {
	resp, err := http.Get(regAddr)
	if err != nil {
		return "", err
	}

	j, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}

	var smk ServerMasterKey
	err = json.Unmarshal(j, &smk)
	if err != nil {
		return "", err
	}

	return smk.MasterPublicKey, nil
}
