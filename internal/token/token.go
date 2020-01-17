package token

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"strings"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

var privateKey *ecdsa.PrivateKey = nil

func init() {
	var err error
	privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		log.Fatal(err)
	}
}

// TODO: make this config
const tokenDuration = time.Hour * 10

type Claim struct {
	UserID string `json:"userId"`
	jwt.Token
}

func New(userID string) (string, error) {
	now := time.Now()

	token := jwt.New()

	token.Set(jwt.IssuedAtKey, now.Unix())
	token.Set(jwt.ExpirationKey, now.Add(tokenDuration).Unix())
	token.Set("UserID", userID)

	signed, err := token.Sign(jwa.ES256, privateKey)

	return string(signed), err
}

func Parse(token string) (*jwt.Token, error) {
	return jwt.Parse(strings.NewReader(token), jwt.WithVerify(jwa.ES256, &privateKey.PublicKey))
}
