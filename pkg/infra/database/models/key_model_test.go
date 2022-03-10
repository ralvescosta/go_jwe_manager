package models

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	valueObjects "jwemanager/pkg/domain/value_objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_KeyModel_ToValueObject(t *testing.T) {
	t.Run("should return value object correctly", func(t *testing.T) {
		priv, _ := rsa.GenerateKey(rand.Reader, 2048)

		sut := KeyModel{
			PubKey: x509.MarshalPKCS1PublicKey(&priv.PublicKey),
			PriKey: x509.MarshalPKCS1PrivateKey(priv),
		}

		vo, err := sut.ToValueObject()

		assert.NoError(t, err)
		assert.IsType(t, valueObjects.Key{}, vo)
	})

	t.Run("should return error when try to convert the pub key", func(t *testing.T) {
		sut := KeyModel{}

		_, err := sut.ToValueObject()

		assert.Error(t, err)
	})

	t.Run("should return error when try to convert the priv key", func(t *testing.T) {
		priv, _ := rsa.GenerateKey(rand.Reader, 2048)

		sut := KeyModel{
			PubKey: x509.MarshalPKCS1PublicKey(&priv.PublicKey),
		}

		_, err := sut.ToValueObject()

		assert.Error(t, err)
	})
}

func Test_ToKeyModel(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		priv, _ := rsa.GenerateKey(rand.Reader, 2048)

		vo := valueObjects.Key{
			ID:     "id",
			UserID: "user_id",
			PriKey: priv,
			PubKey: &priv.PublicKey,
		}

		model := ToKeyModel(vo)

		assert.Equal(t, "id", model.ID)
		assert.Equal(t, "user_id", model.UserID)
	})
}

func Test_KeyModel_MarshalBinary(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := KeyModel{}

		_, err := sut.MarshalBinary()

		assert.NoError(t, err)
	})
}

func Test_KeyModel_UnmarshalBinary(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := KeyModel{}

		b, _ := sut.MarshalBinary()

		err := sut.UnmarshalBinary(b)

		assert.NoError(t, err)
	})
}
