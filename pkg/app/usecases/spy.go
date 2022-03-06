package usecases

import (
	"context"
	valueObjects "jwemanager/pkg/domain/value_objects"

	"github.com/stretchr/testify/mock"
)

//
type CreateKeyUseCaseSpy struct {
	mock.Mock
}

func (pst CreateKeyUseCaseSpy) Execute(ctx context.Context, key valueObjects.Key, timeToExpiration int) (valueObjects.Key, error) {
	args := pst.Called(ctx, key, timeToExpiration)

	return args.Get(0).(valueObjects.Key), args.Error(1)
}

func NewCreateKeyUseCaseSpy() *CreateKeyUseCaseSpy {
	return new(CreateKeyUseCaseSpy)
}

//
type DecryptUseCaseSpy struct {
	mock.Mock
}

func (pst DecryptUseCaseSpy) Decrypt(ctx context.Context, data valueObjects.DecryptValueObject) (valueObjects.DecryptedValueObject, error) {
	args := pst.Called(ctx, data)

	return args.Get(0).(valueObjects.DecryptedValueObject), args.Error(1)
}

func NewDecryptUseCaseSpy() *DecryptUseCaseSpy {
	return new(DecryptUseCaseSpy)
}

//
type EncryptUseCaseSpy struct {
	mock.Mock
}

func (pst EncryptUseCaseSpy) Encrypt(ctx context.Context, data valueObjects.EncryptValueObject) (valueObjects.EncryptedValueObject, error) {
	args := pst.Called(ctx, data)

	return args.Get(0).(valueObjects.EncryptedValueObject), args.Error(1)
}

func NewEncryptUseCaseSpy() *EncryptUseCaseSpy {
	return new(EncryptUseCaseSpy)
}

//
type GetKeyUseCaseSpy struct {
	mock.Mock
}

func (pst GetKeyUseCaseSpy) GetKey(ctx context.Context, userID, keyID string) (valueObjects.Key, error) {
	args := pst.Called(ctx, userID, keyID)

	return args.Get(0).(valueObjects.Key), args.Error(1)
}

func NewGetKeyUseCaseSpy() *GetKeyUseCaseSpy {
	return new(GetKeyUseCaseSpy)
}
