package crypto

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"

	"jwemanager/pkg/app/interfaces"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwe"
)

type crypto struct {
	logger interfaces.ILogger
}

func (pst crypto) Encrypt(pubKey rsa.PrivateKey, data interface{}) ([]byte, error) {
	dataToByte, err := json.Marshal(data)
	if err != nil {
		pst.logger.Error(fmt.Sprintf("[Crypto::Encrypt] Marshal Error: %s", err.Error()))
		return nil, err
	}

	encrypted, err := jwe.Encrypt(dataToByte, jwa.RSA_OAEP_256, pubKey, jwa.A256CBC_HS512, jwa.NoCompress)
	if err != nil {
		pst.logger.Error(fmt.Sprintf("[Crypto::Encrypt] JWE Encrypt Error: %s", err.Error()))
		return nil, err
	}

	return encrypted, nil
}

func (pst crypto) Decrypt(privKey *rsa.PrivateKey, data []byte) (interface{}, error) {

	data, err := jwe.Decrypt(data, jwa.RSA_OAEP_256, privKey)

	pst.logger.Debug(fmt.Sprintf("%v", data))

	if err != nil {
		pst.logger.Error(fmt.Sprintf("[Crypto::Encrypt] JWE Decrypt Error: %s", err.Error()))
		return nil, err
	}

	return nil, nil
}

func NewCrypto() interfaces.ICrypto {
	return crypto{}
}
