package pkgs

import (
	"crypto/rand"
	"fmt"
)

func RandomString(n int) string {
	randBytes := make([]byte, n/2)
	_, err := rand.Read(randBytes)
	if err != nil {
		panic("cannot generator random string")
	}
	return fmt.Sprintf("%x", randBytes)
}
