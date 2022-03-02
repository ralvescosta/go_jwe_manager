package crypto

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"

	"jwemanager/pkg/app/errors"
	"jwemanager/pkg/app/interfaces"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwe"
)

type crypto struct {
	logger interfaces.ILogger
}

var encrypt = jwe.Encrypt
var marshal = json.Marshal
var decrypt = jwe.Decrypt

func (pst crypto) Encrypt(pubKey *rsa.PublicKey, data map[string]interface{}) ([]byte, error) {
	dataToByte, err := marshal(data)
	if err != nil {
		pst.logger.Error(fmt.Sprintf("[Crypto::Encrypt] Marshal Error: %s", err.Error()))
		return nil, errors.NewInternalError(err.Error())
	}

	encrypted, err := encrypt(dataToByte, jwa.RSA_OAEP_256, pubKey, jwa.A256CBC_HS512, jwa.NoCompress)
	if err != nil {
		pst.logger.Error(fmt.Sprintf("[Crypto::Encrypt] JWE Encrypt Error: %s", err.Error()))
		return nil, errors.NewInternalError(err.Error())
	}

	return encrypted, nil
}

func (pst crypto) Decrypt(privKey *rsa.PrivateKey, data []byte) (map[string]interface{}, error) {
	data, err := decrypt(data, jwa.RSA_OAEP_256, privKey)
	if err != nil {
		pst.logger.Error(fmt.Sprintf("[Crypto::Encrypt] JWE Decrypt Error: %s", err.Error()))
		return nil, errors.NewInternalError(err.Error())
	}

	var decrypted = make(map[string]interface{})
	err = json.Unmarshal(data, &decrypted)
	if err != nil {
		pst.logger.Error(fmt.Sprintf("[Crypto::Encrypt] Decrypted Data Unmarshaler Error: %s", err.Error()))
		return nil, errors.NewInternalError(err.Error())
	}

	return decrypted, nil
}

func NewCrypto(logger interfaces.ILogger) interfaces.ICrypto {
	return crypto{logger}
}
