package keyGenerator

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"jwemanager/pkg/app/interfaces"
)

type keyGenerator struct {
	logger interfaces.ILogger
}

// Create RSA key
func (pst keyGenerator) GenerateKey() (*rsa.PrivateKey, error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		pst.logger.Error(fmt.Sprintf("[KeyGenerator::GenerateKey] - Error: %s", err.Error()))
		return nil, err
	}

	return priv, nil
}

// Create a new KeyGenerator instance
func NewKeyGenerator(logger interfaces.ILogger) interfaces.IKeyGenerator {
	return keyGenerator{logger}
}
