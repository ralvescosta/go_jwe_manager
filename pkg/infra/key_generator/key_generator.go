package keyGenerator

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"jwemanager/pkg/app/errors"
	"jwemanager/pkg/app/interfaces"
)

type keyGenerator struct {
	logger interfaces.ILogger
}

var genRSAKey = rsa.GenerateKey

// Create RSA key
func (pst keyGenerator) GenerateKey() (*rsa.PrivateKey, error) {
	priv, err := genRSAKey(rand.Reader, 2048)
	if err != nil {
		pst.logger.Error(fmt.Sprintf("[KeyGenerator::GenerateKey] - Error: %s", err.Error()))
		return nil, errors.NewInternalError(err.Error())
	}

	return priv, nil
}

// Create a new KeyGenerator instance
func NewKeyGenerator(logger interfaces.ILogger) interfaces.IKeyGenerator {
	return keyGenerator{logger}
}
