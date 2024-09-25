package jwt

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

var (
	RsaPub = []byte("-----BEGIN PUBLIC KEY-----\nMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANKv0qZYXEY6/P2TWMDOaeXssQQc1/Us\n3dCKANcl7Lmx7J/sHHcdMsZNRjWttxC/182BkBRIoQ7hetP9BG8wnqkCAwEAAQ==\n-----END PUBLIC KEY-----")
	RsaPri = []byte("-----BEGIN PRIVATE KEY-----\nMIIBUwIBADANBgkqhkiG9w0BAQEFAASCAT0wggE5AgEAAkEA0q/SplhcRjr8/ZNY\nwM5p5eyxBBzX9Szd0IoA1yXsubHsn+wcdx0yxk1GNa23EL/XzYGQFEihDuF60/0E\nbzCeqQIDAQABAkBFCzOIKerLZSdlXjU2si5IGCIGjAFFqpdicOdHmnkSfSEq0O8r\n+Ds/oIqNcJpyPn9ErsJ+23SzNAU9p/eJh//RAiEA6fzX3hGiYugDO0MV941H3xQC\nlF0mRGZpu3Z0IMKlKJ0CIQDmgdJjK39Ts8Mb+yxRI4DqzNp0AH0izcSWtgY5FbfS\nfQIgPfO8FAgHPri/Ykl433qAtQfPRwkCwMl85S2PwbzHjeECIBX0U3eCkxQD0Rd/\nKs9nlEXI0R2vVjvUYV8BY0JYoTN5AiBKT3RwjkeD309dNkzMiIfKs0Qv0/iCN6I9\nUJr6+5OX3g==\n-----END PRIVATE KEY-----")
)

const issuer = "fordisk"

type Jwt interface {
	Parse(token string) (map[string]any, error)
	Generate(payload map[string]interface{}) (string, error)
}

// NewJwter
func NewJwter(pri, pub []byte, expired time.Duration) Jwt {
	prik, err := jwt.ParseRSAPrivateKeyFromPEM(pri)
	if err != nil {
		log.Fatal(err)
	}

	pubk, err := jwt.ParseRSAPublicKeyFromPEM(pub)
	if err != nil {
		log.Fatal(err)
	}

	return &jwter{
		private: prik,
		public:  pubk,
		expired: expired,
	}
}

type jwter struct {
	private *rsa.PrivateKey
	public  *rsa.PublicKey
	expired time.Duration
}

// Parse 用于校验 token 的校验，校验正确则返回 payload 内容，错误则返回报错
func (j *jwter) Parse(tokenString string) (map[string]any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return j.public, nil
	})
	if err != nil {
		return nil, err
	}
	cc, ok := token.Claims.(*CustomerClaims)
	if !ok || cc == nil {
		return nil, fmt.Errorf("token claims type is invalid")
	}
	return cc.Data, nil
}

type CustomerClaims struct {
	Data map[string]interface{}
	jwt.RegisteredClaims
}

// Generate 传入 payload 生成 token 字符串
func (j *jwter) Generate(payload map[string]interface{}) (string, error) {
	now := time.Now()
	mc := CustomerClaims{Data: make(map[string]interface{})}
	mc.ExpiresAt = &jwt.NumericDate{Time: now.Add(j.expired)}
	mc.NotBefore = &jwt.NumericDate{Time: now.Add(-10 * time.Minute)}
	mc.Issuer = issuer
	for k, v := range payload {
		mc.Data[k] = v
	}
	return jwt.NewWithClaims(jwt.SigningMethodRS256, mc).SignedString(j.private)
}
