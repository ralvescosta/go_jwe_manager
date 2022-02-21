package interfaces

import "crypto/rsa"

type ICrypto interface {
	Encrypt(pubKey *rsa.PublicKey, data map[string]interface{}) ([]byte, error)
	Decrypt(privKey *rsa.PrivateKey, data []byte) (map[string]interface{}, error)
}
