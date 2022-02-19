package interfaces

import "crypto/rsa"

type IKeyGenerator interface {
	GenerateKey() (*rsa.PrivateKey, error)
}
