package handlers

import (
	httpServer "jwemanager/pkg/infra/http_server"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Crypto_Encrypt(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewCryptoHandlersSpy()

		req := httpServer.HttpRequest{}
		res := httpServer.HttpResponse{}

		sut.On("Encrypt", req).Return(res)

		result := sut.Encrypt(req)

		assert.Equal(t, res, result)
		sut.AssertExpectations(t)
	})
}

func Test_Crypto_Decrypt(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewCryptoHandlersSpy()

		req := httpServer.HttpRequest{}
		res := httpServer.HttpResponse{}

		sut.On("Decrypt", req).Return(res)

		result := sut.Decrypt(req)

		assert.Equal(t, res, result)
		sut.AssertExpectations(t)
	})
}

func Test_Health_Check(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewHealthHandlerSpy()

		req := httpServer.HttpRequest{}
		res := httpServer.HttpResponse{}

		sut.On("Check", req).Return(res)

		result := sut.Check(req)

		assert.Equal(t, res, result)
		sut.AssertExpectations(t)
	})
}

func Test_Keys_Create(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewKeysHandlersSpy()

		req := httpServer.HttpRequest{}
		res := httpServer.HttpResponse{}

		sut.On("Create", req).Return(res)

		result := sut.Create(req)

		assert.Equal(t, res, result)
		sut.AssertExpectations(t)
	})
}

func Test_Keys_FindOne(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewKeysHandlersSpy()

		req := httpServer.HttpRequest{}
		res := httpServer.HttpResponse{}

		sut.On("FindOne", req).Return(res)

		result := sut.FindOne(req)

		assert.Equal(t, res, result)
		sut.AssertExpectations(t)
	})
}
