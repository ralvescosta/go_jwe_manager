package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewInternalError(t *testing.T) {
	t.Run("should create a internalError correctly", func(t *testing.T) {
		err := NewInternalError("some error")

		assert.Error(t, err)
		assert.IsType(t, InternalError{}, err)
	})
}
