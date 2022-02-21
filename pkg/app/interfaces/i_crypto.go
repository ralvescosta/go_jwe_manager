package interfaces

import "crypto/rsa"

type ICrypto interface {
	Encrypt(pubKey rsa.PrivateKey, data interface{}) ([]byte, error)
	Decrypt(privKey *rsa.PrivateKey, data []byte) (interface{}, error)
}
