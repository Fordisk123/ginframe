package jwt

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	pub := "-----BEGIN PUBLIC KEY-----\nMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANKv0qZYXEY6/P2TWMDOaeXssQQc1/Us\n3dCKANcl7Lmx7J/sHHcdMsZNRjWttxC/182BkBRIoQ7hetP9BG8wnqkCAwEAAQ==\n-----END PUBLIC KEY-----"
	pri := "-----BEGIN PRIVATE KEY-----\nMIIBUwIBADANBgkqhkiG9w0BAQEFAASCAT0wggE5AgEAAkEA0q/SplhcRjr8/ZNY\nwM5p5eyxBBzX9Szd0IoA1yXsubHsn+wcdx0yxk1GNa23EL/XzYGQFEihDuF60/0E\nbzCeqQIDAQABAkBFCzOIKerLZSdlXjU2si5IGCIGjAFFqpdicOdHmnkSfSEq0O8r\n+Ds/oIqNcJpyPn9ErsJ+23SzNAU9p/eJh//RAiEA6fzX3hGiYugDO0MV941H3xQC\nlF0mRGZpu3Z0IMKlKJ0CIQDmgdJjK39Ts8Mb+yxRI4DqzNp0AH0izcSWtgY5FbfS\nfQIgPfO8FAgHPri/Ykl433qAtQfPRwkCwMl85S2PwbzHjeECIBX0U3eCkxQD0Rd/\nKs9nlEXI0R2vVjvUYV8BY0JYoTN5AiBKT3RwjkeD309dNkzMiIfKs0Qv0/iCN6I9\nUJr6+5OX3g==\n-----END PRIVATE KEY-----"

	jwter := NewJwter([]byte(pri), []byte(pub), 24*time.Hour)
	generate, err := jwter.Generate(map[string]interface{}{
		"a": "b",
		"c": 123,
	})
	if err != nil {
		panic(err)
	}
	println(generate)
	if err != nil {
		panic(err)
	}
	pay, err := jwter.Parse("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJEYXRhIjp7ImEiOiJiIiwiYyI6MTIzfSwiaXNzIjoiZm9yZGlzayIsImV4cCI6MTcyNzMxODQxNywibmJmIjoxNzI3MjMxNDE3fQ.M8yZYhl_CdmSDAAofJNUoO4KA7fLetOvakWMxSDx5moHm9ClmB8yxL0w1bLq7i3_ah_CgbbRzmReSTwLCmYpxg")
	if err != nil {
		panic(err)
	}
	for s, a := range pay {
		fmt.Println(s)
		fmt.Println(a)
	}
}
