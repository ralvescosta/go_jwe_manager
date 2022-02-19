package keyGenerator

import (
	"crypto/rand"
	"crypto/rsa"
	"jwemanager/pkg/app/interfaces"
)

type keyGenerator struct{}

// Create RSA key
func (keyGenerator) GenerateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

// Create a new KeyGenerator instance
func NewKeyGenerator() interfaces.IKeyGenerator {
	return keyGenerator{}
}
