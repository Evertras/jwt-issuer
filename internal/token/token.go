package token

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

var privateKey *ecdsa.PrivateKey = nil
var jwkJson string

func init() {
	var err error
	privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		log.Fatal(err)
	}

	key, err := jwk.New(&privateKey.PublicKey)

	if err != nil {
		log.Fatal(err)
	}

	raw, err := json.Marshal(key)

	if err != nil {
		log.Fatal(err)
	}

	jwkJson = string(raw)
}

// TODO: make this config
const tokenDuration = time.Hour * 10

type Claim struct {
	UserID string `json:"userId"`
	jwt.Token
}

func JWKJson() string {
	return jwkJson
}

func New(userID string) (string, error) {
	now := time.Now()

	token := jwt.New()

	token.Set(jwt.IssuedAtKey, now.Unix())
	token.Set(jwt.ExpirationKey, now.Add(tokenDuration).Unix())
	token.Set(jwt.IssuerKey, "evertras/jwt-issuer")
	token.Set(jwt.SubjectKey, userID)
	token.Set("customgreeting", "hello")

	signed, err := token.Sign(jwa.ES256, privateKey)

	return string(signed), err
}

func Parse(token io.Reader) (*jwt.Token, error) {
	return jwt.Parse(token, jwt.WithVerify(jwa.ES256, &privateKey.PublicKey))
}
