# Go JLINC DID

A go client for the JLINC DID Server

Specs:

- https://did-spec.jlinc.org/
- https://w3c.github.io/did-core/

Specifically: https://did-spec.jlinc.org/#6-operations

## Nomenclature

### DID

A DID: `did:jlinc:dnhZF0DuRHgZ0wetY4S7ygsrCQMUzUZoxLxosEmbIYM` is the unique Decentralized Identifier.


### DID Document

A DID Document is a JSON objected sent to and received by the JLINC DID Server. It looks like this

```json
{
  "@context": "https://w3id.org/did/v1",
  "id": "did:jlinc:dnhZF0DuRHgZ0wetY4S7ygsrCQMUzUZoxLxosEmbIYM",
  "created": "2019-01-08T21:12:36.505Z",
  "publicKey": [
    {
      "id": "did:jlinc:dnhZF0DuRHgZ0wetY4S7ygsrCQMUzUZoxLxosEmbIYM#signing",
      "type": "Ed25519VerificationKey2018",
      "controller:": "did:jlinc:dnhZF0DuRHgZ0wetY4S7ygsrCQMUzUZoxLxosEmbIYM",
      "publicKeyBase64": "dnhZF0DuRgsrCQMUzUZoxLHgZ0wetY4S7yxosEmbIYM",
      "publicKeyBase58": "8yTYbTktiu4FS4cda52vkBaidHt9C1KsUk4LbmFEFW42"
    },
    {
      "id": "did:jlinc:dnhZF0DuRHgZ0wetY4S7ygsrCQMUzUZoxLxosEmbIYM#encrypting",
      "type": "X25519KeyAgreementKey2019",
      "controller:": "did:jlinc:dnhZF0DuRHgZ0wetY4S7ygsrCQMUzUZoxLxosEmbIYM",
      "publicKeyBase64": "BK478lOPXtO9J2KsWq_M_opXcVqCiAYd0TWOJcATjX8",
      "publicKeyBase58:": "KGj1nUsanJzBpD1WycQi8FftxQzdk3acn4Sp9LWWD5p"
    },
  ],
}
```

### Private DID data

The private DID data produced by the RegisterDID function is collected into a
struct and marshaled into JSON for persisting. PrivateDidData can also be used
to reify the JSON for reading and writing.

```golang
type PrivateDidData struct {
	SigningPublicKey     string `json:"signingPublicKey"`
	SigningPrivateKey    string `json:"signingPrivateKey"`
	EncryptingPublicKey  string `json:"encryptingPublicKey"`
	EncryptingPrivateKey string `json:"encryptingPrivateKey"`
	RegistrationSecret   string `json:"registrationSecret"`
	DID                  string `json:"did"`
}
```

## Installation and  Usage

```golang
import "github.com/jlinclabs/go-jlinc-did"
// then install it into your app
go mod tidy

// configure a did server
const didServerUrl = "http://localhost:5001"

// register a DID and persist the resulting JSON object
jsonForSaving, err := did.RegisterDID(didServerUrl)

//reconstruct the DID data for usage if desired
var reified did.PrivateDidData
err := json.Unmarshal([]byte(jsonForSaving), &reified)

```

## TODO
  resolve a DID

  supersede a DID
  
  get a DID's history
