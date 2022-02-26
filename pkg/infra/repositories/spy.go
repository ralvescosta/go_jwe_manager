package repositories

import (
	"context"

	valueObjects "jwemanager/pkg/domain/value_objects"

	"github.com/stretchr/testify/mock"
)

type KeyRepositorySpy struct {
	mock.Mock
}

func (pst KeyRepositorySpy) CreateKey(ctx context.Context, key valueObjects.Key, timeToExpiration int) (valueObjects.Key, error) {
	args := pst.Called(ctx, key, timeToExpiration)

	return args.Get(0).(valueObjects.Key), args.Error(1)
}

func (pst KeyRepositorySpy) GetKeyByID(ctx context.Context, userID, keyID string) (valueObjects.Key, error) {
	args := pst.Called(ctx, userID, keyID)

	return args.Get(0).(valueObjects.Key), args.Error(1)
}

func NewKeyRepositorySpy() *KeyRepositorySpy {
	return new(KeyRepositorySpy)
}
